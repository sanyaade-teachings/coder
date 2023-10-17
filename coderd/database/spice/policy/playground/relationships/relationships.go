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
	groupFinance := Group("finance").MemberUser(camilla, ammar, kyle, shark)
	groupCostControl := Group("cost-control").MemberGroup(groupFinance).MemberUser(dean, colin)
	groupEngineers := Group("engineers").MemberUser(ammar, colin, dean, jon, kayla, kira, kyle, steven)
	groupMarketing := Group("marketing").MemberUser(katherine, ammar)
	groupSales := Group("sales").MemberUser(shark, eric)

	// Teams
	teamCompany := Team("company").Platform(platform)
	teamLegal := Team("legal").Platform(platform)
	teamMarketing := Team("marketing").Platform(platform)
	// People who write code
	teamDevelopers := Team("developers").Platform(platform)
	// People who tinker
	teamTechnical := Team("technical").Platform(platform)
	//teamCustomerSuccess := Team("customer-success").Platform(platform) // Customer solutions
	//teamEngineering := Team("engineering").Platform(platform)

	// Nest some teams
	// TODO: This is currently unsupported

	// Assign groups to teams
	teamCompany.MemberGroup(groupEveryone)
	teamLegal.MemberGroup(groupHR, groupFinance)
	teamMarketing.MemberGroup(groupMarketing)
	teamDevelopers.
		Workspace_creatorGroup(groupEngineers).
		// Cost control groups can edit workspaces & delete them
		Workspace_editorGroup(groupCostControl).
		Workspace_deletorGroup(groupCostControl)

	teamTechnical.
		Workspace_creatorGroup(groupEngineers, groupSales).
		// 1 off assignment of a single user.
		Workspace_creatorUser(elliot)

	// Make some resources!
	devTemplate := Template("dev-template").Owner(teamDevelopers)
	devVersion := devTemplate.Version("active")
	var _ = devVersion

	stevenWorkspace := WorkspaceWithDeps("steven-workspace").
		Owner(teamDevelopers).
		ViewerUser(steven).
		EditorUser(steven).
		DeletorUser(steven).
		SelectorUser(steven).
		ConnectorUser(steven)
	devTemplate.Workspace(stevenWorkspace)

}
