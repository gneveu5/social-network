import Axios from "./api"

let posting = (payload) => {
    return Axios.post('/post', payload, {
        headers: {
            "Authorization" : `Bearer ${sessionStorage.getItem("jwt")}`,
            "Content-Type": "multipart/form-data"
        },
    })
}

let commenting = (payload) => {
    return Axios.post('/commenting', payload, {
        headers: {
            "Authorization" : `Bearer ${localStorage.getItem("jwt")}`,
            "Content-Type": "multipart/form-data"
        },
    })
}

let postfetch = (urlarg) => {
    if (urlarg) {
        return Axios.get('/post?id='+urlarg, {
            headers: {
                "Authorization" : `Bearer ${sessionStorage.getItem("jwt")}`
            },
        })
    } else {
        return Axios.get('/post', {
            headers: {
                "Authorization" : `Bearer ${sessionStorage.getItem("jwt")}`
            },
        })
    }
}

let postCommentFetch = (payload) => {
    return Axios.post('/comment', payload, {
        headers: {
            "Authorization" : `Bearer ${sessionStorage.getItem("jwt")}`
        },
    })
}

export const postService = {
    posting, postfetch, postCommentFetch, commenting
}