package relationships

func GenerateRelationships() {
	var (
		platform = Platform("default")

		root = User("root")
		// This is an incomplete list. Using real names to use intuitive groupings.
		ammar     = User("ammar")
		camilla   = User("camilla")
		colin     = User("colin")
		dean      = User("dean")
		elliot    = User("elliot")
		eric      = User("eric")
		jon       = User("jon")
		katherine = User("katherine")
		kayla     = User("kayla")
		kira      = User("kira")
		kyle      = User("kyle")
		shark     = User("shark")
		steven    = User("steven")
	)

	// Add platform roles
	platform.Administrator(root)

	// Groups
	groupEveryone := Group("everyone").MemberWildcard()
	groupHR := Group("hr").MemberUser(camilla)
	groupFinance := Group("finance").MemberUser(ammar, kyle, shark)
	groupCostControl := Group("cost-control").MemberUser(ammar, kyle, dean, colin)
	groupEngineers := Group("engineers").MemberUser(ammar, colin, dean, jon, kayla, kira, kyle, steven)
	groupMarketing := Group("marketing").MemberUser(katherine, ammar)
	groupSales := Group("sales").MemberUser(shark, eric)

	// Teams
	teamCompany := Team("company").
		Platform(platform).
		// Cost control can see all workspaces
		Workspace_viewerGroup(groupCostControl)
	teamLegal := Team("legal").Platform(platform).
		Parent(teamCompany)
	teamMarketing := Team("marketing").Platform(platform).
		Parent(teamCompany)

	// company
	// ├── legal
	// ├── marketing
	// └── engineering
	//      ├── developers
	//      └── technical
	teamEngineering := Team("engineering").Platform(platform).
		Parent(teamCompany)

	// People who write code
	teamDevelopers := Team("developers").Platform(platform).
		Parent(teamEngineering)
	// People who tinker
	teamTechnical := Team("technical").Platform(platform).
		Parent(teamEngineering)

	// Nest some teams
	// TODO: This is currently unsupported

	// Assign groups to teams
	teamCompany.MemberGroup(groupEveryone).
		// Cost control groups can edit workspaces & delete them
		Workspace_editorGroup(groupCostControl).
		Workspace_deletorGroup(groupCostControl)
	teamLegal.MemberGroup(groupHR, groupFinance)
	teamMarketing.MemberGroup(groupMarketing)

	teamDevelopers.
		Workspace_creatorGroup(groupEngineers)

	teamTechnical.
		Workspace_creatorGroup(groupEngineers, groupSales).
		// 1 off assignment of a single user.
		Workspace_creatorUser(elliot)

	// Make some resources!
	devTemplate := Template("dev-template").Owner(teamDevelopers)
	devVersion := devTemplate.Version("active")
	devTemplate.CannotUseBy(teamMarketing)
	var _ = devVersion

	stevenWorkspace := WorkspaceWithDeps("steven-workspace", teamDevelopers, devTemplate).
		ViewerUser(steven).
		EditorUser(steven).
		DeletorUser(steven).
		SelectorUser(steven).
		ConnectorUser(steven)

	// Add some assertions
	stevenWorkspace.
		CanViewBy(steven, ammar, kyle).
		CannotViewBy(camilla, jon)

	// Validations enumerate who can do the given action.
	stevenWorkspace.ValidateView().ValidateSsh().ValidateEdit()
}
