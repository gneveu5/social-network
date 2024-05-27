import ContactList from "./contactList"
import ChatWindow from "./chatWindow"
import { Contact, Conversation } from "./chatClasses"
import { WebsocketContext } from "../_services/websocketProvider"
import { useState, useContext, useEffect, useRef } from "react"
import ReconnectWebsocket from "./reconnectWebsocket"

const ChatMain = () => {
    const conversations = useRef(new Map())

    const [contacts, setContacts] = useState({private:[], group:[]})

    const [currentConversation, setCurrentConversation] = useState(null)

    const chatBottomScroll = useRef(true)

    const [ready, value, send, connect] = useContext(WebsocketContext);

    const updateContacts = () => {
        const updatedContacts = {private:[], group:[]}
        for (const [key, value] of conversations.current) {
            if (value.contact.type === "grp") {
                updatedContacts.group.push(value.contact)
            } else if (value.contact.type === "pvt") {
                updatedContacts.private.push(value.contact)
            }
        }
        setContacts(updatedContacts)
    }

    useEffect(() => {
        if (!ready) {
            setContacts({private:[], group:[]})
            setCurrentConversation(null)
            conversations.current = new Map()
            return
        }

        if (value === null) {
            return
        }

        const data = JSON.parse(value?.data)

        switch (value?.type) {
        case "init-contact":
            initContacts(data)
        break;
        case "connect":
            setConnectedStatus(data.userId, true)
        break;
        case "disconnect":
            setConnectedStatus(data.userId, false)
        break;
        case "pvt-msg":
            addNewMessage(data, "pvt")
        break;
        case "grp-msg":
            addNewMessage(data, "grp")
        break;
        case "pvt-hist":
            addMessageHistory(data, "pvt")
        break;
        case "grp-hist":
            addMessageHistory(data, "grp")
        break;
        case "rmv-pvt":
            removeContact(data, "pvt")
        break;
        case "rmv-grp":
            removeContact(data, "grp")
        break;
        default:
            console.log("message type " + value.type + " not supported yet")
        break;
        
        }
    }, [ready, value])

    const initContacts = (data) => {
        if (data.private) {
            data.private.forEach((newContactData) => {
                initConversation(newContactData, "pvt")
            })
        }
        if (data.group) {
            data.group.forEach((newContactData) => {
                initConversation(newContactData, "grp")
            })
        }
        updateContacts()
    }

    const initConversation = (data, type) => {
        const contact = new Contact(data.contactId, type, data.name)
        if (type === "pvt") {
            contact.setConnected(data.connected)
        }
        const conversation = new Conversation(contact)
        conversation.addMessageHistory(data.history)
        conversations.current.set(type + '-' + data.contactId, conversation)
    }

    const setConnectedStatus = (userId, connectedStatus) => {
        if (conversations.current.has("pvt-" + userId)) {
            conversations.current.get("pvt-" + userId).contact.setConnected(connectedStatus)
            updateContacts()
        }
    }

    const addNewMessage = (data, type) => {
        const conversationId = type + "-" + data.to
        conversations.current.get(conversationId)?.addNewMessage(data)
        if (currentConversation?.id === conversationId) {
            chatBottomScroll.current = true
            setCurrentConversation(conversations.current.get(conversationId).copy)
        } else {
            conversations.current.get(conversationId)?.incrementUnreadMessages()
            updateContacts()
        }
    }

    const addMessageHistory = (data, type) => {
        const conversationId = type + "-" + data.to
        conversations.current.get(conversationId)?.addMessageHistory(data.hist)
        if (currentConversation?.id === conversationId) {
            chatBottomScroll.current = false
            setCurrentConversation(conversations.current.get(conversationId).copy)
        }
    }

    const removeContact = (data, type) => {
        const conversationId = type + "-" + data.id
        conversations.current.delete(conversationId)
        updateContacts()
    }

    const changeConversation = (contact) => {
        if (contact === null) {
            setCurrentConversation(null)
        } else {
            conversations.current.get(contact.type + "-" + contact.id).resetUnreadMessages()
            setCurrentConversation(conversations.current.get(contact.type + "-" + contact.id))
        }
    }

    const chatWindowSendMessage = (data) => {
        if (currentConversation === null || !ready) {
            return
        }
        const webSocketMessage = {type: currentConversation.contact.type + "-" + data.type}
        let payload = {}
        switch (data.type) {
        case "msg":
            payload.txt = data.txt
            payload.to = currentConversation.contact.id
            break;
        case "hist":
            payload.last = data.last
            payload.id = currentConversation.contact.id
            break;
        }
        webSocketMessage.data = JSON.stringify(payload)
        send(JSON.stringify(webSocketMessage))
    }

    return (
        <ReconnectWebsocket containingDivId="contact-window">
            <ContactList 
                contacts={contacts}
                changeConversation={changeConversation}
            />
            <ChatWindow 
                conversation={currentConversation}
                send={chatWindowSendMessage}
                scrollToBottom={chatBottomScroll}
                changeConversation={changeConversation}
            />
        </ReconnectWebsocket> 
    )
}

export default ChatMain