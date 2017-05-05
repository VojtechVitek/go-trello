/*
Copyright 2014 go-trello authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trello

type Action struct {
	client          *Client
	Id              string `json:"id"`
	IdMemberCreator string `json:"idMemberCreator"`
	Data struct {
		DateLastEdited string `json:"dateLastEdited"`
		ListBefore struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"listBefore"`
		ListAfter struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"listAfter"`
		CheckItem struct {
			Id    string `json:"id"`
			State string `json:"state"`
			Name  string `json:"name"`
		} `json:"checkItem"`
		CheckList struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"checklist"`
		List struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"list"`
		TextData struct {
			Emoji struct{} `json:"emoji"`
		} `json:"textData"`
		Board struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
		} `json:"board"`
		Card struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
			IdShort   int    `json:"idShort"`
		} `json:"card"`
		Text string `json:"text"`
	} `json:"data"`
	Type ActionType `json:"type"`
	Date string `json:"date"`
	MemberCreator struct {
		Id         string `json:"id"`
		AvatarHash string `json:"avatarHash"`
		FullName   string `json:"fullName"`
		Initials   string `json:"initials"`
		Username   string `json:"username"`
	} `json:"memberCreator"`
}

type ActionType string

const (
	AddAdminToBoard                   ActionType = "addAdminToBoard"
	AddAdminToOrganization            ActionType = "addAdminToOrganization"
	AddAttachmentToCard               ActionType = "addAttachmentToCard"
	AddBoardsPinnedToMember           ActionType = "addBoardsPinnedToMember"
	AddChecklistToCard                ActionType = "addChecklistToCard"
	AddLabelToCard                    ActionType = "addLabelToCard"
	AddMemberToBoard                  ActionType = "addMemberToBoard"
	AddMemberToCard                   ActionType = "addMemberToCard"
	AddMemberToOrganization           ActionType = "addMemberToOrganization"
	AddToOrganizationBoard            ActionType = "addToOrganizationBoard"
	CommentCard                       ActionType = "commentCard"
	ConvertToCardFromCheckItem        ActionType = "convertToCardFromCheckItem"
	CopyBoard                         ActionType = "copyBoard"
	CopyCard                          ActionType = "copyCard"
	CopyChecklist                     ActionType = "copyChecklist"
	CreateLabel                       ActionType = "createLabel"
	CopyCommentCard                   ActionType = "copyCommentCard"
	CreateBoard                       ActionType = "createBoard"
	CreateBoardInvitation             ActionType = "createBoardInvitation"
	CreateBoardPreference             ActionType = "createBoardPreference"
	CreateCard                        ActionType = "createCard"
	CreateChecklist                   ActionType = "createChecklist"
	CreateList                        ActionType = "createList"
	CreateOrganization                ActionType = "createOrganization"
	CreateOrganizationInvitation      ActionType = "createOrganizationInvitation"
	DeleteAttachmentFromCard          ActionType = "deleteAttachmentFromCard"
	DeleteBoardInvitation             ActionType = "deleteBoardInvitation"
	DeleteCard                        ActionType = "deleteCard"
	DeleteCheckItem                   ActionType = "deleteCheckItem"
	DeleteLabel                       ActionType = "deleteLabel"
	DeleteOrganizationInvitation      ActionType = "deleteOrganizationInvitation"
	DisablePlugin                     ActionType = "disablePlugin"
	DisablePowerUp                    ActionType = "disablePowerUp"
	EmailCard                         ActionType = "emailCard"
	EnablePlugin                      ActionType = "enablePlugin"
	EnablePowerUp                     ActionType = "enablePowerUp"
	MakeAdminOfBoard                  ActionType = "makeAdminOfBoard"
	MakeAdminOfOrganization           ActionType = "makeAdminOfOrganization"
	MakeNormalMemberOfBoard           ActionType = "makeNormalMemberOfBoard"
	MakeNormalMemberOfOrganization    ActionType = "makeNormalMemberOfOrganization"
	MakeObserverOfBoard               ActionType = "makeObserverOfBoard"
	MemberJoinedTrello                ActionType = "memberJoinedTrello"
	MoveCardFromBoard                 ActionType = "moveCardFromBoard"
	MoveCardToBoard                   ActionType = "moveCardToBoard"
	MoveListFromBoard                 ActionType = "moveListFromBoard"
	MoveListToBoard                   ActionType = "moveListToBoard"
	RemoveAdminFromBoard              ActionType = "removeAdminFromBoard"
	RemoveAdminFromOrganization       ActionType = "removeAdminFromOrganization"
	RemoveBoardsPinnedFromMember      ActionType = "removeBoardsPinnedFromMember"
	RemoveChecklistFromCard           ActionType = "removeChecklistFromCard"
	RemoveFromOrganizationBoard       ActionType = "removeFromOrganizationBoard"
	RemoveLabelFromCard               ActionType = "removeLabelFromCard"
	RemoveMemberFromBoard             ActionType = "removeMemberFromBoard"
	RemoveMemberFromCard              ActionType = "removeMemberFromCard"
	RemoveMemberFromOrganization      ActionType = "removeMemberFromOrganization"
	UnconfirmedBoardInvitation        ActionType = "unconfirmedBoardInvitation"
	UnconfirmedOrganizationInvitation ActionType = "unconfirmedOrganizationInvitation"
	UpdateBoard                       ActionType = "updateBoard"
	UpdateCard                        ActionType = "updateCard"
	UpdateCheckItem                   ActionType = "updateCheckItem"
	UpdateCheckItemStateOnCard        ActionType = "updateCheckItemStateOnCard"
	UpdateChecklist                   ActionType = "updateChecklist"
	UpdateLabel                       ActionType = "updateLabel"
	UpdateList                        ActionType = "updateList"
	UpdateMember                      ActionType = "updateMember"
	UpdateOrganization                ActionType = "updateOrganization"
	VoteOnCard                        ActionType = "voteOnCard"
)
