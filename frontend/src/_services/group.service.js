import Axios from "./api"

let groupCreation = (payload) => {
    return Axios.post('/groupcreation', payload)
}

let fetchGroupList = (urlarg) => {
    return Axios.get('/grouplist?id='+urlarg)
}

let groupAskMembership = (payload) => {
    return Axios.post('/groupaskmembership', payload)
}

let groupInviteMembership = (payload) => {
    return Axios.post('/groupinvitemembership', payload)
}

let fetchGroupExist = (urlarg) => {
    return Axios.get('/groupfetchexist?id='+urlarg)
}

let fetchGroupAskMembership = (urlarg) => {
    return Axios.get('/groupfetchaskmembership?id='+urlarg)
}

let fetchGroupMember = (urlarg) => {
    return Axios.get('/groupmember?id='+urlarg)
}

let fetchGroupMembers = (urlarg) => {
    return Axios.get('/groupmembers?id='+urlarg)
}

let eventing = (payload) => {
    return Axios.post('/group', payload)
}

let eventRegister = (payload) => {
    return Axios.post('/event', payload)
}

let eventFetch = (urlarg) => {
    return Axios.get('/group?id='+urlarg)
}

let eventAttendeesFetch = (urlarg) => {
    return Axios.get('/event?id='+urlarg)
}

export const groupService = {
    groupCreation,
    groupAskMembership,
    fetchGroupList,
    groupInviteMembership,
    fetchGroupExist,
    fetchGroupAskMembership,
    fetchGroupMember,
    fetchGroupMembers,
    eventing,
    eventFetch,
    eventAttendeesFetch,
    eventRegister
}