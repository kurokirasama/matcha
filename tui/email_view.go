package tui

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/floatpane/matcha/calendar"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
	"github.com/floatpane/matcha/theme"
	"github.com/floatpane/matcha/view"
)

// ClearKittyGraphics sends the Kitty graphics protocol delete command directly to stdout.
func ClearKittyGraphics() {
	// Delete all images: a=d (action=delete), d=A (delete all)
	os.Stdout.WriteString("\x1b_Ga=d,d=A\x1b\\")
	os.Stdout.Sync()
}

var (
	emailHeaderStyle   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).Padding(0, 1)
	attachmentBoxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).PaddingLeft(2)
)

func (m *EmailView) getHeaderHeight(width int) int {
	var cryptoStatus strings.Builder
	if m.isEncrypted || m.isSMIME || m.isPGPEncrypted || m.isPGP {
		cryptoStatus.WriteString(" [SECURITY]")
	}
	header := fmt.Sprintf("To: %s | From: %s | Subject: %s%s", strings.Join(m.email.To, ", "), m.email.From, m.email.Subject, cryptoStatus.String())
	return lipgloss.Height(emailHeaderStyle.Width(width).Render(header))
}

func (m *EmailView) getAttachmentHeight() int {
	if len(m.email.Attachments) == 0 {
		return 0
	}
	return 1 + len(m.email.Attachments)
}

func (m *EmailView) getCalendarHeight() int {
	if !m.hasCalendarInvite || m.calendarEvent == nil {
		return 0
	}
	return lipgloss.Height(renderCalendarInvite(m.calendarEvent))
}

func (m *EmailView) getHelpText() string {
	if m.focusOnAttachments {
		helpText := "↑/↓: navigate • enter: download • esc/tab: back to email body"
		if m.pluginStatus != "" {
			helpText += " • " + m.pluginStatus
		}
		return helpText
	}
	var shortcuts strings.Builder
	toggleReadKey := config.Keybinds.Email.ToggleRead
	if toggleReadKey == "" {
		toggleReadKey = config.Keybinds.Inbox.ToggleRead
	}
	shortcuts.WriteString(fmt.Sprintf("\uf112 r: reply • \uf064 f: forward • \uea81 d: delete • \uea98 a: archive • %s: read status • \uf435 tab: focus attachments • \ueb06 esc: back to inbox", toggleReadKey))
	if view.ImageProtocolSupported() {
		shortcuts.WriteString(" • \uf03e i: toggle images")
	}
	for _, pk := range m.pluginKeyBindings {
		shortcuts.WriteString(" • ")
		shortcuts.WriteString(pk.Key)
		shortcuts.WriteString(": ")
		shortcuts.WriteString(pk.Description)
	}
	if m.pluginStatus != "" {
		shortcuts.WriteString(" • ")
		shortcuts.WriteString(m.pluginStatus)
	}
	return shortcuts.String()
}

func (m *EmailView) getHelpHeight(width int) int {
	text := m.getHelpText()
	return lipgloss.Height(HelpStyle.Width(width).Render(text))
}

// BodyTransformer, if set, post-processes the rendered email body before it is
// placed in the viewport. main.go wires this up to the plugin manager so that
// plugins registered on the "email_body_render" hook can rewrite, recolor, or
// remove parts of the displayed body.
var BodyTransformer func(body string, email fetcher.Email) string

func applyBodyTransform(body string, email fetcher.Email) string {
	if BodyTransformer == nil {
		return body
	}
	return BodyTransformer(body, email)
}

type EmailView struct {
	viewport           viewport.Model
	email              fetcher.Email
	emailIndex         int
	attachmentCursor   int
	focusOnAttachments bool
	accountID          string
	mailbox            MailboxKind
	folderName         string
	disableImages      bool
	showImages         bool
	isSMIME            bool
	smimeTrusted       bool
	isEncrypted        bool
	isPGP              bool
	pgpTrusted         bool
	isPGPEncrypted     bool
	imagePlacements    []view.ImagePlacement
	pluginStatus       string
	pluginKeyBindings  []PluginKeyBinding
	hasCalendarInvite  bool
	calendarEvent      *calendar.Event
	originalICSData    []byte
	isPreviewMode      bool
	columnOffset       int // horizontal offset for image rendering in split pane
	rowOffset          int // vertical offset for image rendering in horizontal split pane
	totalHeight        int // total height allocated to this view
}

func NewEmailView(email fetcher.Email, emailIndex, width, height int, mailbox MailboxKind, folderName string, disableImages bool) *EmailView {
	isSMIME := false
	smimeTrusted := false
	isEncrypted := false
	isPGP := false
	pgpTrusted := false
	isPGPEncrypted := false
	var filteredAtts []fetcher.Attachment
	var calendarEvent *calendar.Event
	var originalICSData []byte

	for _, att := range email.Attachments {
		if att.Filename == "smime-status.internal" {
			isSMIME = att.IsSMIMESignature || att.IsSMIMEEncrypted
			smimeTrusted = att.SMIMEVerified
			isEncrypted = att.IsSMIMEEncrypted
		} else if att.IsSMIMESignature || att.Filename == "smime.p7s" || att.Filename == "smime.p7m" || strings.HasPrefix(att.MIMEType, "application/pkcs7") {
			if att.IsSMIMESignature && !isSMIME {
				isSMIME = true
				smimeTrusted = att.SMIMEVerified
			}
		} else if att.Filename == "pgp-status.internal" {
			isPGP = att.IsPGPSignature || att.IsPGPEncrypted
			pgpTrusted = att.PGPVerified
			isPGPEncrypted = att.IsPGPEncrypted
		} else if att.IsPGPSignature || att.Filename == "signature.asc" || att.MIMEType == "application/pgp-signature" || att.MIMEType == "application/pgp-encrypted" {
			if att.IsPGPSignature && !isPGP {
				isPGP = true
				pgpTrusted = att.PGPVerified
			}
		} else if att.IsCalendarInvite {
			if len(att.Data) > 0 && calendarEvent == nil {
				if event, err := calendar.ParseICS(att.Data); err == nil {
					calendarEvent = event
					originalICSData = att.Data
				}
			}
		} else {
			filteredAtts = append(filteredAtts, att)
		}
	}
	email.Attachments = filteredAtts

	inlineImages := inlineImagesFromAttachments(email.Attachments)
	showImages := !disableImages

	body, placements, err := view.ProcessBodyWithInline(email.Body, email.BodyMIMEType, inlineImages, H1Style, H2Style, BodyStyle, !showImages)
	if err != nil {
		body = fmt.Sprintf("Error rendering body: %v", err)
	}
	body = applyBodyTransform(body, email)

	m := &EmailView{
		viewport:          viewport.New(),
		email:             email,
		emailIndex:        emailIndex,
		accountID:         email.AccountID,
		mailbox:           mailbox,
		folderName:        folderName,
		disableImages:     disableImages,
		showImages:        showImages,
		isSMIME:           isSMIME,
		smimeTrusted:      smimeTrusted,
		isEncrypted:       isEncrypted,
		isPGP:             isPGP,
		pgpTrusted:        pgpTrusted,
		isPGPEncrypted:    isPGPEncrypted,
		imagePlacements:   placements,
		hasCalendarInvite: calendarEvent != nil,
		calendarEvent:     calendarEvent,
		originalICSData:   originalICSData,
		isPreviewMode:     false,
	}
	
	m.Update(tea.WindowSizeMsg{Width: width, Height: height})
	
	return m
}

func NewEmailViewPreview(email fetcher.Email, folderName string, width, height, colOffset int, disableImages bool) *EmailView {
	ev := NewEmailView(email, 0, width, height, MailboxInbox, folderName, disableImages)
	ev.isPreviewMode = true
	ev.columnOffset = colOffset
	return ev
}

func (m *EmailView) Init() tea.Cmd {
	return nil
}

func (m *EmailView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		kb := config.Keybinds
		if msg.String() == kb.Global.Cancel {
			if m.focusOnAttachments {
				m.focusOnAttachments = false
				return m, nil
			}
			ClearKittyGraphics()
			return m, func() tea.Msg { return BackToMailboxMsg{Mailbox: m.mailbox} }
		}

		if m.focusOnAttachments {
			switch msg.String() {
			case "up", kb.Global.NavUp:
				if len(m.email.Attachments) > 0 {
					m.attachmentCursor = (m.attachmentCursor - 1 + len(m.email.Attachments)) % len(m.email.Attachments)
				}
				return m, nil
			case "down", kb.Global.NavDown:
				if len(m.email.Attachments) > 0 {
					m.attachmentCursor = (m.attachmentCursor + 1) % len(m.email.Attachments)
				}
				return m, nil
			case "enter":
				if len(m.email.Attachments) > 0 {
					selected := m.email.Attachments[m.attachmentCursor]
					idx := m.emailIndex
					accountID := m.accountID
					return m, func() tea.Msg {
						return DownloadAttachmentMsg{
							Index:     idx,
							Filename:  selected.Filename,
							PartID:    selected.PartID,
							Data:      selected.Data,
							AccountID: accountID,
							Mailbox:   m.mailbox,
						}
					}
				}
			case kb.Email.FocusAttachments:
				m.focusOnAttachments = false
			}
		} else {
			switch msg.String() {
			case kb.Email.ToggleImages:
				if view.ImageProtocolSupported() {
					m.showImages = !m.showImages
					ClearKittyGraphics()
					inlineImages := inlineImagesFromAttachments(m.email.Attachments)
					body, placements, err := view.ProcessBodyWithInline(m.email.Body, m.email.BodyMIMEType, inlineImages, H1Style, H2Style, BodyStyle, !m.showImages)
					if err != nil {
						body = fmt.Sprintf("Error rendering body: %v", err)
					}
					body = applyBodyTransform(body, m.email)
					m.imagePlacements = placements
					wrapped := wrapBodyToWidth(body, m.viewport.Width())
					m.viewport.SetContent(wrapped)
					return m, nil
				}
			case kb.Email.Reply:
				ClearKittyGraphics()
				return m, func() tea.Msg { return ReplyToEmailMsg{Email: m.email} }
			case kb.Email.Forward:
				ClearKittyGraphics()
				return m, func() tea.Msg { return ForwardEmailMsg{Email: m.email} }
			case kb.Email.Delete:
				accountID := m.accountID
				uid := m.email.UID
				ClearKittyGraphics()
				return m, func() tea.Msg {
					return DeleteEmailMsg{UID: uid, AccountID: accountID, Mailbox: m.mailbox}
				}
			case kb.Email.Archive:
				accountID := m.accountID
				uid := m.email.UID
				ClearKittyGraphics()
				return m, func() tea.Msg {
					return ArchiveEmailMsg{UID: uid, AccountID: accountID, Mailbox: m.mailbox}
				}
			case kb.Email.RsvpAccept, kb.Email.RsvpDecline, kb.Email.RsvpTentative:
				if m.hasCalendarInvite && m.calendarEvent != nil {
					var response string
					switch msg.String() {
					case kb.Email.RsvpAccept:
						response = "ACCEPTED"
					case kb.Email.RsvpDecline:
						response = "DECLINED"
					case kb.Email.RsvpTentative:
						response = "TENTATIVE"
					}
					return m, func() tea.Msg {
						return SendRSVPMsg{
							OriginalICS: m.originalICSData,
							Event:       m.calendarEvent,
							Response:    response,
							AccountID:   m.accountID,
							InReplyTo:   m.email.MessageID,
							References:  m.email.References,
						}
					}
				}
			case kb.Email.FocusAttachments:
				if len(m.email.Attachments) > 0 {
					m.focusOnAttachments = true
				}
			case kb.Inbox.ToggleRead:
				key := kb.Email.ToggleRead
				if key == "" {
					key = kb.Inbox.ToggleRead
				}
				if msg.String() == key {
					uid := m.email.UID
					accountID := m.accountID
					isRead := m.email.IsRead
					ClearKittyGraphics()
					return m, tea.Batch(func() tea.Msg {
						if isRead {
							return MarkEmailAsUnreadMsg{UID: uid, AccountID: accountID, FolderName: m.mailbox.folderName(m.folderName)}
						}
						return MarkEmailAsReadMsg{UID: uid, AccountID: accountID, FolderName: m.mailbox.folderName(m.folderName)}
					}, func() tea.Msg { return BackToMailboxMsg{Mailbox: m.mailbox} })
				}
			}
		}
	case tea.WindowSizeMsg:
		m.totalHeight = msg.Height
		m.viewport.SetWidth(msg.Width)
		headerHeight := m.getHeaderHeight(msg.Width)
		calendarHeight := m.getCalendarHeight()
		attachmentHeight := m.getAttachmentHeight()
		helpHeight := m.getHelpHeight(msg.Width)
		
		// Budget exactly for: Header\n, Calendar\n, Body, Attachments\n, Help
		// Spacer count = (Header exists? 1 : 0) + (Calendar exists? 1 : 0) + (Attachments exists? 1 : 0) + 1 (before Help)
		spacers := 1 // Help spacer
		if headerHeight > 0 { spacers++ }
		if calendarHeight > 0 { spacers++ }
		if attachmentHeight > 0 { spacers++ }
		
		vh := msg.Height - headerHeight - calendarHeight - attachmentHeight - helpHeight - spacers
		if vh < 1 { vh = 1 }
		m.viewport.SetHeight(vh)

		ClearKittyGraphics()
		inlineImages := inlineImagesFromAttachments(m.email.Attachments)
		body, placements, err := view.ProcessBodyWithInline(m.email.Body, m.email.BodyMIMEType, inlineImages, H1Style, H2Style, BodyStyle, !m.showImages)
		if err != nil {
			body = fmt.Sprintf("Error rendering body: %v", err)
		}
		body = applyBodyTransform(body, m.email)
		m.imagePlacements = placements
		wrapped := wrapBodyToWidth(body, m.viewport.Width())
		m.viewport.SetContent(wrapped)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *EmailView) View() tea.View {
	os.Stdout.WriteString("\x1b_Ga=d,d=a\x1b\\")
	os.Stdout.Sync()

	var cryptoStatus strings.Builder
	if m.isEncrypted {
		cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Accent).Render(" [S/MIME: 🔒 Encrypted]"))
	} else if m.isSMIME {
		if m.smimeTrusted {
			cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Accent).Render(" [S/MIME: ✅ Trusted]"))
		} else {
			cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Danger).Render(" [S/MIME: ❌ Untrusted]"))
		}
	}
	if m.isPGPEncrypted {
		cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Accent).Render(" [PGP: 🔒 Encrypted]"))
	} else if m.isPGP {
		if m.pgpTrusted {
			cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Accent).Render(" [PGP: ✅ Verified]"))
		} else {
			cryptoStatus.WriteString(lipgloss.NewStyle().Foreground(theme.ActiveTheme.Danger).Render(" [PGP: ⚠️ Unverified]"))
		}
	}

	header := fmt.Sprintf("To: %s | From: %s | Subject: %s%s", strings.Join(m.email.To, ", "), m.email.From, m.email.Subject, cryptoStatus.String())
	styledHeader := emailHeaderStyle.Width(m.viewport.Width()).Render(header)

	if m.showImages && len(m.imagePlacements) > 0 {
		headerLines := lipgloss.Height(styledHeader) + 1
		yOffset := m.viewport.YOffset()
		vpHeight := m.viewport.Height()
		for i := range m.imagePlacements {
			p := &m.imagePlacements[i]
			if p.Line >= yOffset && p.Line < yOffset+vpHeight {
				screenRow := m.rowOffset + headerLines + (p.Line - yOffset)
				if m.columnOffset > 0 {
					view.RenderImageToStdout(p, screenRow, m.columnOffset+1)
				} else {
					view.RenderImageToStdout(p, screenRow)
				}
			}
		}
	}

	var b strings.Builder
	b.WriteString(styledHeader)
	b.WriteString("\n")
	if m.hasCalendarInvite && m.calendarEvent != nil {
		b.WriteString(renderCalendarInvite(m.calendarEvent))
		b.WriteString("\n")
	}
	b.WriteString(m.viewport.View())

	if len(m.email.Attachments) > 0 {
		b.WriteString("\n")
		var attB strings.Builder
		attB.WriteString("Attachments:\n")
		for i, attachment := range m.email.Attachments {
			cursor := "  "
			style := itemStyle
			if m.focusOnAttachments && i == m.attachmentCursor {
				cursor = "> "
				style = selectedItemStyle
			}
			attB.WriteString(style.Render(fmt.Sprintf("%s%s", cursor, attachment.Filename)))
			if i < len(m.email.Attachments)-1 {
				attB.WriteString("\n")
			}
		}
		b.WriteString(attachmentBoxStyle.Render(attB.String()))
	}

	b.WriteString("\n")
	helpText := m.getHelpText()
	b.WriteString(HelpStyle.Width(m.viewport.Width()).Render(helpText))

	content := strings.TrimSuffix(b.String(), "\n")
	
	currentHeight := lipgloss.Height(content)
	if currentHeight < m.totalHeight {
		content += strings.Repeat("\n", m.totalHeight - currentHeight)
	}
	
	return tea.NewView(content)
}

func (m *EmailView) GetAccountID() string { return m.accountID }
func (m *EmailView) SetPluginStatus(status string) { m.pluginStatus = status }
func (m *EmailView) SetPluginKeyBindings(bindings []PluginKeyBinding) { m.pluginKeyBindings = bindings }

func inlineImagesFromAttachments(atts []fetcher.Attachment) []view.InlineImage {
	var imgs []view.InlineImage
	for _, att := range atts {
		if !att.Inline || len(att.Data) == 0 || att.ContentID == "" { continue }
		imgs = append(imgs, view.InlineImage{CID: att.ContentID, Base64: base64.StdEncoding.EncodeToString(att.Data)})
	}
	return imgs
}

func wrapBodyToWidth(body string, width int) string { return BodyStyle.Width(width).Render(body) }
func (m *EmailView) GetEmail() fetcher.Email { return m.email }
func (m *EmailView) SetOffsets(row, col int) { m.rowOffset = row; m.columnOffset = col }

func renderCalendarInvite(event *calendar.Event) string {
	if event == nil { return "" }
	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(theme.ActiveTheme.Accent).Padding(1, 2)
	var b strings.Builder
	b.WriteString("📅 Meeting Invite\n\n")
	b.WriteString(fmt.Sprintf("Title:    %s\n", event.Summary))
	b.WriteString(fmt.Sprintf("When:     %s\n", formatEventTime(event.Start, event.End)))
	if event.Location != "" { b.WriteString(fmt.Sprintf("Where:    %s\n", event.Location)) }
	b.WriteString(fmt.Sprintf("Organizer: %s\n", event.Organizer))
	if event.Description != "" { b.WriteString(fmt.Sprintf("\n%s\n", truncateString(event.Description, 100))) }
	b.WriteString("\n")
	b.WriteString(lipgloss.NewStyle().Italic(true).Render("Press 1:Accept  2:Decline  3:Tentative"))
	return style.Render(b.String())
}

func formatEventTime(start, end time.Time) string {
	start = start.Local(); end = end.Local()
	if start.Format("2006-01-02") == end.Format("2006-01-02") {
		return fmt.Sprintf("%s, %s - %s", start.Format("Mon Jan 2, 2006"), start.Format("3:04 PM"), end.Format("3:04 PM"))
	}
	return fmt.Sprintf("%s - %s", start.Format("Mon Jan 2 3:04 PM"), end.Format("Mon Jan 2 3:04 PM"))
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen { return s }
	return s[:maxLen] + "..."
}
