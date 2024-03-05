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
	// ├── everyone
	// ├── hr
	// ├── finance
	// ├── cost-control
	// ├── engineers
	// ├── marketing
	// └── sales

	// └── cost-control
	//      ├── developers
	//      └── technical

	groupEveryone := Group("everyone").MemberWildcard()
	groupHR := Group("hr").MemberUser(camilla)
	groupFinance := Group("finance").MemberUser(ammar, kyle, shark)
	groupCostControl := Group("cost-control").MemberUser(ammar, kyle, dean, colin)
	groupEngineers := Group("engineers").MemberUser(ammar, colin, dean, jon, kayla, kira, kyle, steven)
	groupMarketing := Group("marketing").MemberUser(katherine, ammar)
	groupSales := Group("sales").MemberUser(shark, eric)
	var _ = groupEveryone

	// Organizations
	companyOrganization := Organization("company").
		Platform(platform).
		// Primitive membership
		MemberGroup(groupHR, groupFinance, groupMarketing, groupSales, groupEngineers, groupCostControl).
		// Cost control can see all workspaces
		Workspace_viewerGroup(groupCostControl).
		Workspace_editorGroup(groupCostControl).
		// Engineers & Sales can create workspaces
		Workspace_creatorGroup(groupEngineers, groupSales).
		// External person who can also creaet workspaces
		Workspace_creatorUser(elliot).
		Template_viewerGroup(groupCostControl, groupEngineers, groupSales)

	// Make some resources!
	devTemplate := Template("dev-template").Organization(companyOrganization)
	devVersion := devTemplate.Version("active")
	devTemplate.CannotUseBy(elliot, groupHR)
	var _ = devVersion

	// Steven will create a workspace.
	companyOrganization.CanCreate_workspaceBy(steven) // Perm check
	stevenWorkspace := WorkspaceWithDeps("steven-workspace", companyOrganization, devTemplate).
		ViewerUser(steven).
		EditorUser(steven).
		DeletorUser(steven).
		SelectorUser(steven).
		ConnectorUser(steven)

	// Add some assertions
	stevenWorkspace.
		CanViewBy(steven, ammar, kyle).
		CannotViewBy(camilla, jon)

	// The workspace can be edited by cost control group via teamCompany
	stevenWorkspace.
		CanEditBy(dean).
		// But cloud cost cannot exec into the workspace.
		CannotSshBy(dean)

	// Validations enumerate who can do the given action.
	stevenWorkspace.ValidateView().ValidateSsh().ValidateEdit()
}

// createWorkspace
//   - actor: The user creating the workspace. This user will be assigned as the owner.
//   - team: The team the workspace is being created for.
//   - template: The template version the workspace is being created from.
//   - provisioner: (in prod this might be tags??) The provisioner to provision the workspace.
//
// Creating a workspace is the process of a Team creating a workspace and assigning
// a user permissions.
// Perm checks:
//   - Can a user create a workspace for a given team?
//   - Can the team provision the workspace with the template?
//   - Can the team use the selected provisioner to provision the workspace? (TODO, rethink this)
func testCreateWorkspace(actor *ObjUser, team *ObjOrganization, version *ObjTemplate_version) {

}
