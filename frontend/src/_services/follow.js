import Axios from "./api";

//liste des followers
let followersApi = (credentials) => {
    return Axios.get('/follower?query='+ credentials)
}

let followersNotInGroup = (groupId) => {
    return Axios.get('/followernotingroup?id='+ groupId)
}

//liste des following
let followingApi = (credentials) => {
    return Axios.get('/following?query='+ credentials)
}
// follow un profil public
let reqfollow = (credentials) => {
    return Axios.post('/reqfollow?query='+ credentials)
}
// unfollow un profil
let requnfollow = (credentials) => {
    return Axios.post('/requnfollow?query='+ credentials)
}
// demande de follow à un profil priver
let reqfollowprivate = (credentials) => {
    return Axios.post('/reqfollow?query='+ credentials)
}
// annuler la demande follow à un profil privé
let reqcancelfollowprivate = (credentials) => {
    return Axios.post('/reqcancelfollowprivate?query='+ credentials)
}

export const followService = {
    followersApi, followersNotInGroup, followingApi, reqfollow, requnfollow, reqfollowprivate, reqcancelfollowprivate
}
