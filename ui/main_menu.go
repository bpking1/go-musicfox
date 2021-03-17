package ui

import (
    tea "github.com/anhoder/bubbletea"
    "github.com/anhoder/go-musicfox/constants"
    "github.com/muesli/termenv"
    "math"
    "strings"
    "time"
    "unicode/utf8"
)

type mainMenuModel struct {
    doubleColumn    bool // 是否双列显示

    menuTitle     	     string // 菜单标题
    menuTitleStartRow    int    // 菜单标题开始行
    menuTitleStartColumn int    // 菜单标题开始列

    menuStartRow    int // 菜单开始行
    menuStartColumn int // 菜单开始列
    menuBottomRow   int // 菜单最底部所在行

    menuCurPage	 int // 菜单当前页
    menuPageSize int // 菜单每页大小

    menuList      []string // 菜单列表
    selectedIndex int	   // 当前选中的菜单index
}

// update main ui
func updateMainUI(msg tea.Msg, m *neteaseModel) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        return keyMsgHandle(msg, m)

    case tickMainUIMsg:
        // every second update ui
        return m, tickMainUI(time.Second)

    case tea.WindowSizeMsg:
        m.doubleColumn = msg.Width>=80

        // 菜单开始行、列
        m.menuStartRow = msg.Height/3
        if m.doubleColumn {
            m.menuStartColumn = (msg.Width-60)/2
            m.menuBottomRow = m.menuStartRow+int(math.Ceil(float64(m.menuPageSize)/2))-1
        } else {
            m.menuStartColumn = (msg.Width-20)/2
            m.menuBottomRow = m.menuStartRow+m.menuPageSize-1
        }

        // 菜单标题开始行、列
        m.menuTitleStartColumn = m.menuStartColumn
        if constants.MainShowTitle && m.menuStartRow > 2 {
            if m.menuStartRow > 4 {
                m.menuTitleStartRow = m.menuStartRow - 3
            }
        } else if !constants.MainShowTitle && m.menuStartRow > 1 {
            if m.menuStartRow > 3 {
                m.menuTitleStartColumn = m.menuStartRow - 3
            }
        }

    }

    return m, nil
}

// get main ui view
func mainUIView(m *neteaseModel) string {
    if m.WindowWidth <= 0 || m.WindowHeight <= 0 {
        return ""
    }

    var builder strings.Builder

    // 距离顶部的行数
    var top int

    // title
    if constants.MainShowTitle {
        builder.WriteString(titleView(m, &top))
    }

    // menu title
    builder.WriteString(menuTitleView(m, &top))

    // menu list


    return builder.String()
}

// title view
func titleView(m *neteaseModel, top *int) string {
    var titleBuilder strings.Builder
    titleLen := utf8.RuneCountInString(constants.AppName)+2
    prefixLen := (m.WindowWidth-titleLen)/2
    suffixLen := m.WindowWidth-prefixLen-titleLen
    titleBuilder.WriteString(strings.Repeat("─", prefixLen))
    titleBuilder.WriteString(" ")
    titleBuilder.WriteString(strings.ToUpper(constants.AppName))
    titleBuilder.WriteString(" ")
    titleBuilder.WriteString(strings.Repeat("─", suffixLen))

    *top++

    return SetFgStyle(titleBuilder.String(), primaryColor)
}

// menu title
func menuTitleView(m *neteaseModel, top *int) string {
    var menuTitleBuilder strings.Builder
    var title = m.menuTitle
    if len(m.menuTitle) > 50 {
        menuTitleRunes := []rune(m.menuTitle)
        title = string(menuTitleRunes[:50])
    }

    menuTitleBuilder.WriteString(strings.Repeat("\n", m.menuTitleStartRow - *top))
    menuTitleBuilder.WriteString(strings.Repeat(" ", m.menuTitleStartColumn))
    menuTitleBuilder.WriteString(SetFgStyle(title, termenv.ANSIGreen))

    *top = m.menuTitleStartRow

    return menuTitleBuilder.String()
}

func menuListView(m *neteaseModel, top *int) string {
    var menuListBuilder strings.Builder
    menus := getCurPageMenus(m)
    var lines int
    if m.doubleColumn {
        lines = int(math.Ceil(float64(len(menus))/2))
    } else {
        lines = len(menus)
    }

    menuListBuilder.WriteString(strings.Repeat("\n", m.menuStartRow-*top))

    for i := 0; i < lines; i++ {
        if m.doubleColumn {

        }
    }

    return menuListBuilder.String()
}

// 获取当前页的菜单
func getCurPageMenus(m *neteaseModel) []string {
    start := (m.menuCurPage - 1) * m.menuPageSize
    end := int(math.Min(float64(len(m.menuList)), float64(m.menuCurPage*m.menuPageSize)))

    return m.menuList[start:end]
}

// key handle
func keyMsgHandle(msg tea.KeyMsg, m *neteaseModel) (tea.Model, tea.Cmd) {
    return m, nil
}