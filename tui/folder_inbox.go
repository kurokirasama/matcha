package tui

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

const sidebarWidth = 25

var (
	sidebarStyle = lipgloss.NewStyle().
			Width(sidebarWidth).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			PaddingRight(1).
			PaddingLeft(1)

	sidebarTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("42")).
				Bold(true).
				PaddingBottom(1)

	folderStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1)

	activeFolderStyle = lipgloss.NewStyle().
				PaddingLeft(1).
				PaddingRight(1).
				Background(lipgloss.Color("42")).
				Foreground(lipgloss.Color("#000000")).
				Bold(true)

	moveOverlayStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#25A065")).
				Padding(1, 2)

	moveOverlayTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("42")).
				Bold(true).
				PaddingBottom(1)

	moveItemStyle = lipgloss.NewStyle().
			PaddingLeft(1)

	moveSelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(1).
				Foreground(lipgloss.Color("42")).
				Bold(true)

	inboxPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			PaddingRight(1)

	previewPaneStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				PaddingLeft(1)

	focusedBorderColor   = lipgloss.Color("42")
	unfocusedBorderColor = lipgloss.Color("240")
)

type PaneType int

const (
	FocusInbox PaneType = iota
	FocusPreview
)

// FolderInbox combines a folder sidebar with an email list.
type FolderInbox struct {
	folders         []string
	unread          map[string]int
	activeFolderIdx int
	currentFolder   string
	inbox           *Inbox
	accounts        []config.Account
	width           int
	height          int
	isLoadingEmails bool

	// Move-to-folder overlay state
	movingEmail      bool
	moveTargetIdx    int
	moveUID          uint32   // Legacy: single UID
	moveUIDs         []uint32 // Batch: multiple UIDs
	moveAccountID    string
	moveSourceFolder string

	// Image rendering preference, propagated from config.
	disableImages bool

	hideSidebar bool

	// Layout orientation
	layout config.LayoutMode

	// Whether the quick toggle (Shift+L) is enabled in settings
	enableQuickToggle bool

	// Whether the split is currently active (toggled via Shift+L)
	splitActive bool

	// Split pane state
	previewPane        *EmailView
	previewedUID       uint32
	previewedAccountID string
	previewSearchEmail *fetcher.Email
	focusedPane        PaneType
}

func (m *FolderInbox) GetUnreadCountsCopy() map[string]int {
	if m.unread == nil {
		return make(map[string]int)
	}
	result := make(map[string]int)
	maps.Copy(result, m.unread)
	return result
}

// sortFolders sorts folder names with INBOX always first, then alphabetically.
func sortFolders(folders []string) []string {
	sorted := make([]string, len(folders))
	copy(sorted, folders)
	sort.SliceStable(sorted, func(i, j int) bool {
		iUpper := strings.ToUpper(sorted[i])
		jUpper := strings.ToUpper(sorted[j])
		if iUpper == keyINBOX {
			return true
		}
		if jUpper == keyINBOX {
			return false
		}
		return sorted[i] < sorted[j]
	})
	return sorted
}

// SetDateFormat propagates the configured date layout to the inner inbox.
func (m *FolderInbox) SetDateFormat(layout string) {
	if m.inbox != nil {
		m.inbox.SetDateFormat(layout)
	}
}

// SetDetailedDates propagates the detailed date display toggle.
func (m *FolderInbox) SetDetailedDates(enabled bool) {
	if m.inbox != nil {
		m.inbox.SetDetailedDates(enabled)
	}
}

// SetLayout updates the split layout mode.
func (m *FolderInbox) SetLayout(layout config.LayoutMode) {
	m.layout = layout
}

// SetEnableQuickToggle updates whether the quick toggle is enabled.
func (m *FolderInbox) SetEnableQuickToggle(enabled bool) {
	m.enableQuickToggle = enabled
}

// SetSplitActive updates whether the split is active.
func (m *FolderInbox) SetSplitActive(active bool) {
	m.splitActive = active
}

// SetDefaultThreaded propagates the global default threading toggle.
func (m *FolderInbox) SetDefaultThreaded(v bool) {
	if m.inbox != nil {
		m.inbox.SetDefaultThreaded(v)
	}
}

// SetDisableImages propagates the global image-display preference.
func (m *FolderInbox) SetDisableImages(v bool) {
	m.disableImages = v
}

// NewFolderInbox creates a new FolderInbox.
func NewFolderInbox(folders []string, accounts []config.Account) *FolderInbox {
	folders = sortFolders(folders)
	currentFolder := keyINBOX
	if len(folders) > 0 {
		currentFolder = folders[0]
	}

	inbox := NewInbox(nil, accounts)
	inbox.SetFolderName(currentFolder)

	fi := &FolderInbox{
		folders:         folders,
		activeFolderIdx: 0,
		currentFolder:   currentFolder,
		inbox:           inbox,
		accounts:        accounts,
		splitActive:     true,
	}
	fi.updateHelpKeys()
	return fi
}

func (m *FolderInbox) Init() tea.Cmd {
	return nil
}

func (m *FolderInbox) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:gocyclo
	if m.movingEmail {
		return m.updateMoveOverlay(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if m.inbox.list.FilterState() == list.Filtering {
			break
		}
		if m.inbox.searchOverlay != nil {
			break
		}
		kb := config.Keybinds
		if m.previewPane != nil && m.focusedPane == FocusPreview {
			s := msg.String()
			if s != kb.Folder.FocusInbox && s != kb.Folder.FocusPreview && s != kb.Global.Cancel && s != "q" &&
				s != kb.Inbox.ToggleSidebar {
				var cmd tea.Cmd
				var newP tea.Model
				newP, cmd = m.previewPane.Update(msg)
				m.previewPane = newP.(*EmailView)
				return m, cmd
			}
		}
		switch msg.String() {
		case kb.Folder.FocusPreview:
			if m.previewPane != nil && m.focusedPane == FocusInbox {
				m.focusedPane = FocusPreview
				return m, nil
			}
		case kb.Folder.FocusInbox:
			if m.previewPane != nil && m.focusedPane == FocusPreview {
				m.focusedPane = FocusInbox
				return m, nil
			}
		case kb.Folder.NextFolder:
			m.activeFolderIdx = (m.activeFolderIdx + 1) % len(m.folders)
			return m, m.switchFolder()
		case kb.Folder.PrevFolder:
			m.activeFolderIdx = (m.activeFolderIdx - 1 + len(m.folders)) % len(m.folders)
			return m, m.switchFolder()
		case kb.Global.Cancel:
			if m.previewPane != nil {
				m.closeSplitPreview()
				return m, nil
			}
		case kb.Folder.Move:
			if m.inbox.visualMode && len(m.inbox.selectedUIDs) > 0 {
				m.movingEmail = true
				m.moveTargetIdx = 0
				m.moveUIDs = make([]uint32, len(m.inbox.selectionOrder))
				copy(m.moveUIDs, m.inbox.selectionOrder)
				m.moveAccountID = ""
				for _, acctID := range m.inbox.selectedUIDs {
					m.moveAccountID = acctID
					break
				}
				m.moveSourceFolder = m.currentFolder
				return m, nil
			} else {
				selectedItem, ok := m.inbox.list.SelectedItem().(item)
				if ok {
					m.movingEmail = true
					m.moveTargetIdx = 0
					m.moveUID = selectedItem.uid
					m.moveUIDs = []uint32{selectedItem.uid}
					m.moveAccountID = selectedItem.accountID
					m.moveSourceFolder = m.currentFolder
					return m, nil
				}
			}
		case kb.Inbox.ToggleSidebar:
			m.hideSidebar = !m.hideSidebar
			return m.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
		case kb.Folder.ToggleLayout:
			if m.layout != config.LayoutHorizontal && m.enableQuickToggle {
				m.splitActive = !m.splitActive
				return m, tea.Batch(
					func() tea.Msg { return ToggleLayoutMsg{} },
					func() tea.Msg { return tea.WindowSizeMsg{Width: m.width, Height: m.height} },
				)
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.previewPane != nil || m.previewedUID != 0 {
			inboxWidth := m.calculateInboxWidth()
			previewWidth := m.calculatePreviewWidth()
			ih := m.calculateInboxHeight()
			ph := m.calculatePreviewHeight()

			iw_inner := inboxWidth - 2
			if m.layout == config.LayoutHorizontal {
				iw_inner = inboxWidth
			}
			m.inbox.SetSize(iw_inner, ih)

			if m.previewPane != nil {
				pw_inner := previewWidth - 2
				if m.layout == config.LayoutHorizontal {
					pw_inner = previewWidth
				}
				previewMsg := tea.WindowSizeMsg{Width: pw_inner, Height: ph}
				newModel, _ := m.previewPane.Update(previewMsg)
				m.previewPane = newModel.(*EmailView)

				sw := sidebarWidth
				if m.hideSidebar {
					sw = 0
				}
				if m.layout == config.LayoutHorizontal {
					m.previewPane.SetOffsets(ih+1, sw)
				} else {
					m.previewPane.SetOffsets(0, sw+inboxWidth)
				}
			}
		} else {
			sw := sidebarWidth
			if m.hideSidebar {
				sw = 0
			}
			inboxWidth := msg.Width - sw - 3
			if inboxWidth < 20 {
				inboxWidth = 20
			}
			h := msg.Height
			if m.layout == config.LayoutOff && !m.splitActive {
				h = msg.Height / 2
			}
			m.inbox.SetSize(inboxWidth, h)
		}
		return m, nil

	case FolderEmailsFetchedMsg:
		if msg.FolderName != m.currentFolder {
			return m, nil
		}
		m.isLoadingEmails = false
		m.inbox.isFetching = false
		m.inbox.isRefreshing = false
		m.inbox.SetEmails(msg.Emails, m.accounts)
		m.inbox.SetFolderName(msg.FolderName)
		return m, nil

	case FolderEmailsAppendedMsg:
		if msg.FolderName != m.currentFolder {
			return m, nil
		}
		m.inbox.isFetching = false
		m.inbox.list.Title = m.inbox.getTitle()
		if len(msg.Emails) == 0 {
			if m.inbox.noMoreByAccount == nil {
				m.inbox.noMoreByAccount = make(map[string]bool)
			}
			m.inbox.noMoreByAccount[msg.AccountID] = true
			return m, nil
		}
		for _, email := range msg.Emails {
			m.inbox.emailsByAccount[email.AccountID] = append(m.inbox.emailsByAccount[email.AccountID], email)
			m.inbox.allEmails = append(m.inbox.allEmails, email)
		}
		m.inbox.emailCountByAcct[msg.AccountID] = len(m.inbox.emailsByAccount[msg.AccountID])
		m.inbox.updateList()
		return m, nil

	case EmailMovedMsg:
		m.inbox.RemoveEmail(msg.UID, msg.AccountID)
		if msg.UID == m.previewedUID {
			m.closeSplitPreview()
		}
		return m, nil

	case UpdatePreviewMsg:
		m.previewedUID = msg.UID
		m.previewedAccountID = msg.AccountID
		return m, nil

	case PreviewBodyFetchedMsg:
		if msg.UID != m.previewedUID {
			return m, nil
		}
		email := m.findEmailByUID(msg.UID, msg.AccountID)
		if email == nil {
			return m, nil
		}
		email.Body = msg.Body
		email.BodyMIMEType = msg.BodyMIMEType
		email.Attachments = msg.Attachments

		previewWidth := m.calculatePreviewWidth()
		previewHeight := m.calculatePreviewHeight()
		inboxWidth := m.calculateInboxWidth()
		colOffset := sidebarWidth + 2 + inboxWidth
		if m.hideSidebar {
			colOffset -= sidebarWidth
		}
		m.previewPane = NewEmailViewPreview(*email, m.currentFolder, previewWidth, previewHeight, colOffset, m.disableImages)
		return m, nil
	}

	var cmd tea.Cmd
	var newI tea.Model
	newI, cmd = m.inbox.Update(msg)
	m.inbox = newI.(*Inbox)

	if cmd != nil {
		return m, m.wrapInboxCmd(cmd)
	}

	return m, cmd
}

func (m *FolderInbox) wrapInboxCmd(cmd tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		msg := cmd()
		switch inner := msg.(type) {
		case FetchMoreEmailsMsg:
			return FetchFolderMoreEmailsMsg{
				Offset:     inner.Offset,
				AccountID:  inner.AccountID,
				FolderName: m.currentFolder,
				Limit:      inner.Limit,
			}
		case RequestRefreshMsg:
			inner.FolderName = m.currentFolder
			return inner
		case SearchRequestedMsg:
			inner.FolderName = m.currentFolder
			return inner
		}
		return msg
	}
}

func (m *FolderInbox) updateMoveOverlay(msg tea.Msg) (tea.Model, tea.Cmd) {
	kb := config.Keybinds
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case kb.Global.Cancel:
			m.movingEmail = false
			return m, nil
		case "up", kb.Global.NavUp:
			m.moveTargetIdx = (m.moveTargetIdx - 1 + len(m.moveFolderChoices())) % len(m.moveFolderChoices())
			return m, nil
		case "down", kb.Global.NavDown:
			m.moveTargetIdx = (m.moveTargetIdx + 1) % len(m.moveFolderChoices())
			return m, nil
		case "enter":
			choices := m.moveFolderChoices()
			if len(choices) > 0 && m.moveTargetIdx < len(choices) {
				destFolder := choices[m.moveTargetIdx]
				m.movingEmail = false
				if len(m.moveUIDs) > 1 {
					uids := m.moveUIDs
					m.moveUIDs = nil
					m.inbox.visualMode = false
					m.inbox.selectedUIDs = make(map[uint32]string)
					m.inbox.selectionOrder = []uint32{}
					m.inbox.updateListTitle()
					return m, func() tea.Msg {
						return BatchMoveEmailsMsg{
							UIDs:         uids,
							AccountID:    m.moveAccountID,
							SourceFolder: m.moveSourceFolder,
							DestFolder:   destFolder,
						}
					}
				} else {
					return m, func() tea.Msg {
						return MoveEmailToFolderMsg{
							UID:          m.moveUID,
							AccountID:    m.moveAccountID,
							SourceFolder: m.moveSourceFolder,
							DestFolder:   destFolder,
						}
					}
				}
			}
		}
	}
	return m, nil
}

func (m *FolderInbox) moveFolderChoices() []string {
	var choices []string
	for _, f := range m.folders {
		if f != m.currentFolder {
			choices = append(choices, f)
		}
	}
	return choices
}

func (m *FolderInbox) switchFolder() tea.Cmd {
	if m.activeFolderIdx >= 0 && m.activeFolderIdx < len(m.folders) {
		prevFolder := m.currentFolder
		m.currentFolder = m.folders[m.activeFolderIdx]
		m.isLoadingEmails = true
		m.inbox.SetFolderName(m.currentFolder)
		m.inbox.SetEmails(nil, m.accounts)
		folder := m.currentFolder
		return func() tea.Msg {
			return SwitchFolderMsg{FolderName: folder, PreviousFolder: prevFolder}
		}
	}
	return nil
}

func (m *FolderInbox) View() tea.View {
	var content string
	var sidebar string
	if !m.hideSidebar {
		sidebar = m.renderSidebar()
	}

	if m.previewPane != nil {
		inboxPane := m.renderInboxPane()
		previewPane := m.renderPreviewPane()

		if m.layout == config.LayoutHorizontal {
			// JoinVertical adds 1 line. Ensure total height matches m.height.
			content = lipgloss.JoinVertical(lipgloss.Left, inboxPane, previewPane)
			if m.hideSidebar {
				return tea.NewView(strings.TrimSuffix(content, "\n"))
			}
			content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
		} else {
			if m.splitActive {
				if m.hideSidebar {
					content = lipgloss.JoinHorizontal(lipgloss.Top, inboxPane, previewPane)
				} else {
					content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, inboxPane, previewPane)
				}
			} else {
				if m.hideSidebar {
					content = inboxPane
				} else {
					content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, inboxPane)
				}
			}
		}
	} else if m.previewedUID != 0 {
		inboxPane := m.renderInboxPane()
		emptyPreview := m.renderEmptyPreview()
		if m.layout == config.LayoutHorizontal {
			content = lipgloss.JoinVertical(lipgloss.Left, inboxPane, emptyPreview)
			if m.hideSidebar {
				return tea.NewView(strings.TrimSuffix(content, "\n"))
			}
			content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
		} else {
			if m.splitActive {
				if m.hideSidebar {
					content = lipgloss.JoinHorizontal(lipgloss.Top, inboxPane, emptyPreview)
				} else {
					content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, inboxPane, emptyPreview)
				}
			} else {
				if m.hideSidebar {
					content = inboxPane
				} else {
					content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, inboxPane)
				}
			}
		}
	} else {
		inboxView := m.inbox.View().Content
		if m.hideSidebar {
			content = inboxView
		} else {
			content = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, inboxView)
		}
	}

	if m.movingEmail {
		content = m.renderWithMoveOverlay(content)
	}

	return tea.NewView(strings.TrimSuffix(content, "\n"))
}

func (m *FolderInbox) renderSidebar() string {
	var b strings.Builder
	title := t("folder_inbox.folders_title")
	if len(m.accounts) > 0 {
		acc := m.accounts[0]
		if acc.Name != "" {
			title = acc.Name
		} else if acc.FetchEmail != "" {
			title = acc.FetchEmail
		}
	}
	b.WriteString(sidebarTitleStyle.Render(title))
	b.WriteString("\n")
	for i, folder := range m.folders {
		displayName := m.formatFolderName(folder)
		unread := m.unread[folder]

		var tab string
		if unread > 0 {
			tab = fmt.Sprintf("%s (%d)", displayName, unread)
		} else {
			tab = displayName
		}

		if i == m.activeFolderIdx {
			b.WriteString(activeFolderStyle.Width(sidebarWidth - 4).Render(tab))
		} else {
			b.WriteString(folderStyle.Render(tab))
		}
		if i < len(m.folders)-1 {
			b.WriteString("\n")
		}
	}
	return sidebarStyle.Height(m.height).Render(b.String())
}

func (m *FolderInbox) formatFolderName(name string) string {
	name = strings.TrimPrefix(name, "[Gmail]/")
	name = strings.TrimPrefix(name, "[Google Mail]/")
	maxLen := sidebarWidth - 5
	if len(name) > maxLen {
		name = name[:maxLen-1] + "\u2026"
	}
	return name
}

func (m *FolderInbox) renderWithMoveOverlay(content string) string {
	choices := m.moveFolderChoices()
	if len(choices) == 0 {
		return content
	}
	var b strings.Builder
	title := t("folder_inbox.move_to_folder")
	if len(m.moveUIDs) > 1 {
		title = tn("folder_inbox.move_multiple", len(m.moveUIDs), map[string]interface{}{"count": len(m.moveUIDs)})
	}
	b.WriteString(moveOverlayTitleStyle.Render(title))
	b.WriteString("\n")
	for i, folder := range choices {
		displayName := m.formatFolderName(folder)
		if i == m.moveTargetIdx {
			b.WriteString(moveSelectedItemStyle.Render("> " + displayName))
		} else {
			b.WriteString(moveItemStyle.Render("  " + displayName))
		}
		if i < len(choices)-1 {
			b.WriteString("\n")
		}
	}
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render(t("folder_inbox.help")))
	overlay := moveOverlayStyle.Render(b.String())
	contentLines := strings.Split(content, "\n")
	overlayLines := strings.Split(overlay, "\n")
	startRow := (len(contentLines) - len(overlayLines)) / 2
	if startRow < 0 {
		startRow = 0
	}
	for i, overlayLine := range overlayLines {
		row := startRow + i
		if row >= len(contentLines) {
			break
		}
		contentLines[row] = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, overlayLine)
	}
	return strings.Join(contentLines, "\n")
}

func (m *FolderInbox) SetFolders(folders []string) {
	m.folders = sortFolders(folders)
	found := false
	for i, f := range m.folders {
		if f == m.currentFolder {
			m.activeFolderIdx = i
			found = true
			break
		}
	}
	if !found && len(m.folders) > 0 {
		m.activeFolderIdx = 0
		m.currentFolder = m.folders[0]
	}
}

func (m *FolderInbox) SetUnreadCounts(counts map[string]int) {
	m.unread = counts
}

func (m *FolderInbox) DecrementUnreadCount(folder string) {
	if m.unread == nil {
		return
	}
	if m.unread[folder] > 0 {
		m.unread[folder]--
	}
}

func (m *FolderInbox) SetEmails(emails []fetcher.Email, accounts []config.Account) {
	m.accounts = accounts
	m.inbox.SetEmails(emails, accounts)
}

func (m *FolderInbox) GetCurrentFolder() string { return m.currentFolder }
func (m *FolderInbox) HasSplitPreview() bool { return m.previewPane != nil }
func (m *FolderInbox) GetInbox() *Inbox { return m.inbox }
func (m *FolderInbox) GetAccounts() []config.Account { return m.accounts }
func (m *FolderInbox) RemoveEmail(uid uint32, accountID string) { m.inbox.RemoveEmail(uid, accountID) }

func (m *FolderInbox) updateHelpKeys() {
	bindings := []key.Binding{
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next folder")),
		key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev folder")),
		key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "move")),
	}
	if m.previewPane != nil || m.previewedUID != 0 {
		bindings = append(bindings,
			key.NewBinding(key.WithKeys("]"), key.WithHelp("]/[", "switch pane")),
			key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "close preview")),
		)
	}
	m.inbox.extraShortHelpKeys = bindings
}

func (m *FolderInbox) SetLoadingEmails(loading bool) {
	m.isLoadingEmails = loading
	m.inbox.isFetching = loading
}

func (m *FolderInbox) SetRefreshing(refreshing bool) { m.inbox.isRefreshing = refreshing }
func (m *FolderInbox) GetFolders() []string { return m.folders }

func (m *FolderInbox) renderInboxPane() string {
	inboxWidth := m.calculateInboxWidth()
	inboxHeight := m.calculateInboxHeight()
	borderColor := unfocusedBorderColor
	if m.focusedPane == FocusInbox {
		borderColor = focusedBorderColor
	}
	paneStyle := inboxPaneStyle.BorderForeground(borderColor).Width(inboxWidth).Height(inboxHeight)
	contentHeight := inboxHeight
	if m.layout == config.LayoutHorizontal {
		paneStyle = paneStyle.Border(lipgloss.NormalBorder(), false, false, true, false).PaddingRight(0)
		contentHeight--
	}
	m.inbox.SetSize(inboxWidth-2, contentHeight)
	return paneStyle.Render(m.inbox.View().Content)
}

func (m *FolderInbox) renderPreviewPane() string {
	if m.previewPane == nil {
		return m.renderEmptyPreview()
	}
	previewWidth := m.calculatePreviewWidth()
	previewHeight := m.calculatePreviewHeight()
	borderColor := unfocusedBorderColor
	if m.focusedPane == FocusPreview {
		borderColor = focusedBorderColor
	}
	paneStyle := previewPaneStyle.BorderForeground(borderColor).Width(previewWidth).Height(previewHeight)
	if m.layout == config.LayoutHorizontal {
		paneStyle = paneStyle.Border(lipgloss.NormalBorder(), false, false, false, false).PaddingLeft(0)
	}
	return paneStyle.Render(m.previewPane.View().Content)
}

func (m *FolderInbox) renderEmptyPreview() string {
	previewWidth := m.calculatePreviewWidth()
	previewHeight := m.calculatePreviewHeight()
	emptyStyle := lipgloss.NewStyle().Width(previewWidth).Height(previewHeight).Align(lipgloss.Center, lipgloss.Center).Foreground(lipgloss.Color("240"))
	if m.layout == config.LayoutHorizontal {
		emptyStyle = emptyStyle.Border(lipgloss.NormalBorder(), false, false, false, false).PaddingLeft(0)
	}
	return emptyStyle.Render("Loading...")
}

func (m *FolderInbox) OpenSplitPreview(uid uint32, accountID string, email *fetcher.Email) {
	m.previewPane = nil
	m.previewedUID = uid
	m.previewedAccountID = accountID
	m.previewSearchEmail = email
	m.focusedPane = FocusPreview
	inboxWidth := m.calculateInboxWidth()
	m.inbox.SetSize(inboxWidth-2, m.calculateInboxHeight())
	m.updateHelpKeys()
}

func (m *FolderInbox) closeSplitPreview() {
	ClearKittyGraphics()
	m.previewPane = nil
	m.previewedUID = 0
	m.previewedAccountID = ""
	m.previewSearchEmail = nil
	m.focusedPane = FocusInbox
	sw := sidebarWidth
	if m.hideSidebar {
		sw = 0
	}
	m.inbox.SetSize(m.width-sw-3, m.height)
	m.updateHelpKeys()
}

func (m *FolderInbox) findEmailByUID(uid uint32, accountID string) *fetcher.Email {
	for i := range m.inbox.allEmails {
		if m.inbox.allEmails[i].UID == uid && m.inbox.allEmails[i].AccountID == accountID {
			return &m.inbox.allEmails[i]
		}
	}
	if m.previewSearchEmail != nil && m.previewSearchEmail.UID == uid && m.previewSearchEmail.AccountID == accountID {
		return m.previewSearchEmail
	}
	return nil
}

func (m *FolderInbox) calculatePreviewWidth() int {
	sw := sidebarWidth
	if m.hideSidebar {
		sw = 0
	}
	if m.layout == config.LayoutHorizontal {
		return m.width - sw - 2
	}
	remainingWidth := m.width - sw - 4
	inboxWidth := int(float64(remainingWidth) * 0.4)
	if inboxWidth < 30 {
		inboxWidth = 30
	}
	previewWidth := remainingWidth - inboxWidth
	if previewWidth < 40 {
		previewWidth = 40
	}
	return previewWidth
}

func (m *FolderInbox) calculatePreviewHeight() int {
	if m.layout == config.LayoutHorizontal {
		// JoinVertical in View() adds 1 line.
		h := m.height - m.calculateInboxHeight() - 1
		if h < 5 {
			h = 5
		}
		return h
	}
	return m.height
}

func (m *FolderInbox) calculateInboxWidth() int {
	sw := sidebarWidth
	if m.hideSidebar {
		sw = 0
	}
	if m.layout == config.LayoutHorizontal {
		return m.width - sw - 2
	}
	if m.layout == config.LayoutVertical && !m.splitActive {
		return m.width - sw - 2
	}
	remainingWidth := m.width - sw - 4
	inboxWidth := int(float64(remainingWidth) * 0.4)
	if inboxWidth < 30 {
		inboxWidth = 30
	}
	return inboxWidth
}

func (m *FolderInbox) calculateInboxHeight() int {
	if m.layout == config.LayoutHorizontal {
		h := m.height / 2
		if h < 5 {
			h = 5
		}
		return h
	}
	if m.layout == config.LayoutOff && !m.splitActive {
		return m.height / 2
	}
	return m.height
}
