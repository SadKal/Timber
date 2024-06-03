

export default {
  "meta": {},
  "id": "_default",
  "_regex": {},
  "_paramKeys": {},
  "file": {
    "path": "src/routes/_module.svelte",
    "dir": "src/routes",
    "base": "_module.svelte",
    "ext": ".svelte",
    "name": "_module"
  },
  "asyncModule": () => import('../src/routes/_module.svelte'),
  "rootName": "default",
  "routifyDir": import.meta.url,
  "children": [
    {
      "meta": {},
      "id": "_default_chat",
      "_regex": {},
      "_paramKeys": {},
      "name": "chat",
      "file": {
        "path": "src/routes/chat/_module.svelte",
        "dir": "src/routes/chat",
        "base": "_module.svelte",
        "ext": ".svelte",
        "name": "_module"
      },
      "asyncModule": () => import('../src/routes/chat/_module.svelte'),
      "children": [
        {
          "meta": {},
          "id": "_default_chat_ChatLanding_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "ChatLanding",
          "file": {
            "path": "src/routes/chat/ChatLanding.svelte",
            "dir": "src/routes/chat",
            "base": "ChatLanding.svelte",
            "ext": ".svelte",
            "name": "ChatLanding"
          },
          "asyncModule": () => import('../src/routes/chat/ChatLanding.svelte'),
          "children": []
        },
        {
          "meta": {},
          "id": "_default_chat_ChatMain_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "ChatMain",
          "file": {
            "path": "src/routes/chat/ChatMain.svelte",
            "dir": "src/routes/chat",
            "base": "ChatMain.svelte",
            "ext": ".svelte",
            "name": "ChatMain"
          },
          "asyncModule": () => import('../src/routes/chat/ChatMain.svelte'),
          "children": []
        },
        {
          "meta": {
            "dynamic": true
          },
          "id": "_default_chat__chatID__svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "[chatID]",
          "file": {
            "path": "src/routes/chat/[chatID].svelte",
            "dir": "src/routes/chat",
            "base": "[chatID].svelte",
            "ext": ".svelte",
            "name": "[chatID]"
          },
          "asyncModule": () => import('../src/routes/chat/[chatID].svelte'),
          "children": []
        },
        {
          "meta": {},
          "id": "_default_chat__Dashboard",
          "_regex": {},
          "_paramKeys": {},
          "name": "_Dashboard",
          "module": false,
          "file": {
            "path": "src/routes/chat/_Dashboard",
            "dir": "src/routes/chat",
            "base": "_Dashboard",
            "ext": "",
            "name": "_Dashboard"
          },
          "children": [
            {
              "meta": {},
              "id": "_default_chat__Dashboard_ChatDashboard_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "ChatDashboard",
              "file": {
                "path": "src/routes/chat/_Dashboard/ChatDashboard.svelte",
                "dir": "src/routes/chat/_Dashboard",
                "base": "ChatDashboard.svelte",
                "ext": ".svelte",
                "name": "ChatDashboard"
              },
              "asyncModule": () => import('../src/routes/chat/_Dashboard/ChatDashboard.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__Dashboard_ChatThumb_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "ChatThumb",
              "file": {
                "path": "src/routes/chat/_Dashboard/ChatThumb.svelte",
                "dir": "src/routes/chat/_Dashboard",
                "base": "ChatThumb.svelte",
                "ext": ".svelte",
                "name": "ChatThumb"
              },
              "asyncModule": () => import('../src/routes/chat/_Dashboard/ChatThumb.svelte'),
              "children": []
            }
          ]
        },
        {
          "meta": {},
          "id": "_default_chat__Invitations",
          "_regex": {},
          "_paramKeys": {},
          "name": "_Invitations",
          "module": false,
          "file": {
            "path": "src/routes/chat/_Invitations",
            "dir": "src/routes/chat",
            "base": "_Invitations",
            "ext": "",
            "name": "_Invitations"
          },
          "children": [
            {
              "meta": {},
              "id": "_default_chat__Invitations_Invitation_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "Invitation",
              "file": {
                "path": "src/routes/chat/_Invitations/Invitation.svelte",
                "dir": "src/routes/chat/_Invitations",
                "base": "Invitation.svelte",
                "ext": ".svelte",
                "name": "Invitation"
              },
              "asyncModule": () => import('../src/routes/chat/_Invitations/Invitation.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__Invitations_InvitationsList_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "InvitationsList",
              "file": {
                "path": "src/routes/chat/_Invitations/InvitationsList.svelte",
                "dir": "src/routes/chat/_Invitations",
                "base": "InvitationsList.svelte",
                "ext": ".svelte",
                "name": "InvitationsList"
              },
              "asyncModule": () => import('../src/routes/chat/_Invitations/InvitationsList.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__Invitations_InvitationsModal_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "InvitationsModal",
              "file": {
                "path": "src/routes/chat/_Invitations/InvitationsModal.svelte",
                "dir": "src/routes/chat/_Invitations",
                "base": "InvitationsModal.svelte",
                "ext": ".svelte",
                "name": "InvitationsModal"
              },
              "asyncModule": () => import('../src/routes/chat/_Invitations/InvitationsModal.svelte'),
              "children": []
            }
          ]
        },
        {
          "meta": {},
          "id": "_default_chat__Messages",
          "_regex": {},
          "_paramKeys": {},
          "name": "_Messages",
          "module": false,
          "file": {
            "path": "src/routes/chat/_Messages",
            "dir": "src/routes/chat",
            "base": "_Messages",
            "ext": "",
            "name": "_Messages"
          },
          "children": [
            {
              "meta": {},
              "id": "_default_chat__Messages_LoadMoreMessages_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "LoadMoreMessages",
              "file": {
                "path": "src/routes/chat/_Messages/LoadMoreMessages.svelte",
                "dir": "src/routes/chat/_Messages",
                "base": "LoadMoreMessages.svelte",
                "ext": ".svelte",
                "name": "LoadMoreMessages"
              },
              "asyncModule": () => import('../src/routes/chat/_Messages/LoadMoreMessages.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__Messages_Message_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "Message",
              "file": {
                "path": "src/routes/chat/_Messages/Message.svelte",
                "dir": "src/routes/chat/_Messages",
                "base": "Message.svelte",
                "ext": ".svelte",
                "name": "Message"
              },
              "asyncModule": () => import('../src/routes/chat/_Messages/Message.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__Messages__MessageContext",
              "_regex": {},
              "_paramKeys": {},
              "name": "_MessageContext",
              "module": false,
              "file": {
                "path": "src/routes/chat/_Messages/_MessageContext",
                "dir": "src/routes/chat/_Messages",
                "base": "_MessageContext",
                "ext": "",
                "name": "_MessageContext"
              },
              "children": [
                {
                  "meta": {},
                  "id": "_default_chat__Messages__MessageContext_ContextMenu_svelte",
                  "_regex": {},
                  "_paramKeys": {},
                  "name": "ContextMenu",
                  "file": {
                    "path": "src/routes/chat/_Messages/_MessageContext/ContextMenu.svelte",
                    "dir": "src/routes/chat/_Messages/_MessageContext",
                    "base": "ContextMenu.svelte",
                    "ext": ".svelte",
                    "name": "ContextMenu"
                  },
                  "asyncModule": () => import('../src/routes/chat/_Messages/_MessageContext/ContextMenu.svelte'),
                  "children": []
                },
                {
                  "meta": {},
                  "id": "_default_chat__Messages__MessageContext_EditMessageForm_svelte",
                  "_regex": {},
                  "_paramKeys": {},
                  "name": "EditMessageForm",
                  "file": {
                    "path": "src/routes/chat/_Messages/_MessageContext/EditMessageForm.svelte",
                    "dir": "src/routes/chat/_Messages/_MessageContext",
                    "base": "EditMessageForm.svelte",
                    "ext": ".svelte",
                    "name": "EditMessageForm"
                  },
                  "asyncModule": () => import('../src/routes/chat/_Messages/_MessageContext/EditMessageForm.svelte'),
                  "children": []
                },
                {
                  "meta": {},
                  "id": "_default_chat__Messages__MessageContext_EditMessageModal_svelte",
                  "_regex": {},
                  "_paramKeys": {},
                  "name": "EditMessageModal",
                  "file": {
                    "path": "src/routes/chat/_Messages/_MessageContext/EditMessageModal.svelte",
                    "dir": "src/routes/chat/_Messages/_MessageContext",
                    "base": "EditMessageModal.svelte",
                    "ext": ".svelte",
                    "name": "EditMessageModal"
                  },
                  "asyncModule": () => import('../src/routes/chat/_Messages/_MessageContext/EditMessageModal.svelte'),
                  "children": []
                }
              ]
            }
          ]
        },
        {
          "meta": {},
          "id": "_default_chat__NewChat",
          "_regex": {},
          "_paramKeys": {},
          "name": "_NewChat",
          "module": false,
          "file": {
            "path": "src/routes/chat/_NewChat",
            "dir": "src/routes/chat",
            "base": "_NewChat",
            "ext": "",
            "name": "_NewChat"
          },
          "children": [
            {
              "meta": {},
              "id": "_default_chat__NewChat_NewChat_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "NewChat",
              "file": {
                "path": "src/routes/chat/_NewChat/NewChat.svelte",
                "dir": "src/routes/chat/_NewChat",
                "base": "NewChat.svelte",
                "ext": ".svelte",
                "name": "NewChat"
              },
              "asyncModule": () => import('../src/routes/chat/_NewChat/NewChat.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__NewChat_NewChatModal_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "NewChatModal",
              "file": {
                "path": "src/routes/chat/_NewChat/NewChatModal.svelte",
                "dir": "src/routes/chat/_NewChat",
                "base": "NewChatModal.svelte",
                "ext": ".svelte",
                "name": "NewChatModal"
              },
              "asyncModule": () => import('../src/routes/chat/_NewChat/NewChatModal.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__NewChat_NewChatUserList_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "NewChatUserList",
              "file": {
                "path": "src/routes/chat/_NewChat/NewChatUserList.svelte",
                "dir": "src/routes/chat/_NewChat",
                "base": "NewChatUserList.svelte",
                "ext": ".svelte",
                "name": "NewChatUserList"
              },
              "asyncModule": () => import('../src/routes/chat/_NewChat/NewChatUserList.svelte'),
              "children": []
            },
            {
              "meta": {},
              "id": "_default_chat__NewChat_UserResult_svelte",
              "_regex": {},
              "_paramKeys": {},
              "name": "UserResult",
              "file": {
                "path": "src/routes/chat/_NewChat/UserResult.svelte",
                "dir": "src/routes/chat/_NewChat",
                "base": "UserResult.svelte",
                "ext": ".svelte",
                "name": "UserResult"
              },
              "asyncModule": () => import('../src/routes/chat/_NewChat/UserResult.svelte'),
              "children": []
            }
          ]
        },
        {
          "meta": {},
          "id": "_default_chat_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/chat/index.svelte",
            "dir": "src/routes/chat",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/chat/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_index_svelte",
      "_regex": {},
      "_paramKeys": {},
      "name": "index",
      "file": {
        "path": "src/routes/index.svelte",
        "dir": "src/routes",
        "base": "index.svelte",
        "ext": ".svelte",
        "name": "index"
      },
      "asyncModule": () => import('../src/routes/index.svelte'),
      "children": []
    },
    {
      "meta": {},
      "id": "_default_login",
      "_regex": {},
      "_paramKeys": {},
      "name": "login",
      "module": false,
      "file": {
        "path": "src/routes/login",
        "dir": "src/routes",
        "base": "login",
        "ext": "",
        "name": "login"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_login_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/login/index.svelte",
            "dir": "src/routes/login",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/login/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_register",
      "_regex": {},
      "_paramKeys": {},
      "name": "register",
      "module": false,
      "file": {
        "path": "src/routes/register",
        "dir": "src/routes",
        "base": "register",
        "ext": "",
        "name": "register"
      },
      "children": [
        {
          "meta": {},
          "id": "_default_register_index_svelte",
          "_regex": {},
          "_paramKeys": {},
          "name": "index",
          "file": {
            "path": "src/routes/register/index.svelte",
            "dir": "src/routes/register",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/register/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {
        "dynamic": true,
        "dynamicSpread": true
      },
      "_regex": {},
      "_paramKeys": {},
      "name": "[...404]",
      "file": {
        "path": ".routify/components/[...404].svelte",
        "dir": ".routify/components",
        "base": "[...404].svelte",
        "ext": ".svelte",
        "name": "[...404]"
      },
      "asyncModule": () => import('./components/[...404].svelte'),
      "children": []
    }
  ]
}