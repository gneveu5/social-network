import Axios from "./api";



// searchbar get all users and groups
let searchbarleftmenu = (credentials) => {
    return Axios.get('/searchbar?query='+ credentials)
}

export const searchService = {
    searchbarleftmenu

}

