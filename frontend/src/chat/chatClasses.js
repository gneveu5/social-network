// Classes used to represent 

export class Message {
    constructor(id, senderName, content, creationTimeStamp) {
        this.id = id
        this.senderName = senderName
        this.content = content
        this.creationTimeStamp = creationTimeStamp
    }
}

export class Contact {
    constructor(id, type, name) {
        this.id = id
        this.type = type
        this.name = name
        this.unreadMessages = 0
    }

    setConnected(connected) {
        this.connected = connected
    }

    incrementUnreadMessages() {
        this.unreadMessages++
    }

    resetUnreadMessages() {
        this.unreadMessages = 0
    }
}

export class Conversation {
    constructor(contact) {
        this.contact = contact
        this.messageList = []
    }

    get copy() {
        const copy = new Conversation(this.contact)
        copy.contact.unreadMessages = this.contact.unreadMessages
        copy.messageList = [...this.messageList]

        return copy
    }

    get id () {
        return this.contact.type + "-" + this.contact.id
    }

    addNewMessage(message) {
        const newMessage = new Message(message.id, message.from, message.txt, message.date)
        this.messageList.push(newMessage)
    }

    addMessageHistory(history) {
        const oldMessages = []
        history?.forEach((message) => {
            const oldMessage = new Message(message.id, message.from, message.txt, message.date)
            oldMessages.push(oldMessage)
        })
        this.messageList = [...oldMessages, ...this.messageList]
    }

    incrementUnreadMessages() {
        this.contact.unreadMessages++
    }

    resetUnreadMessages() {
        this.contact.unreadMessages = 0
    }
}