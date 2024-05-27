import { postService } from "../../_services/post.service"
import { followService } from "../../_services/follow"
import { useState, useEffect } from "react"

export const PostForm = ({ handlePostSubmitReturn }) => {

    const urlParams = window.location.href

    const [postTitle, setPostTitle] = useState("")
    const [postContent, setPostContent] = useState("")
    const [postPrivacy, setPostPrivacy] = useState("")
    const [picture, setPicture] = useState("")
    const [followers, setFollowers] = useState([])
    const [errors, setErrors] = useState({})

    const handlePostSubmit = async (e) => {

        e.preventDefault()
        let payload
        const formData = new FormData()

        if (urlParams.split('/')[3] === "group") {
            payload = {
                title: postTitle,
                body: postContent,
                groupId: urlParams.split('/')[4],
            }
        } else {
            payload = {
                title: postTitle,
                body: postContent,
                viewStatus: postPrivacy,
                followers : followers,
            }
        }

        formData.append("json", JSON.stringify(payload))
        formData.append("file", picture)

        await postService.posting(formData)
        .then((r) => {
            let tmp = r.data
            for (let i in tmp) {
                tmp[i]["Opened"] = "See comments"
            }
            handlePostSubmitReturn(tmp)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    function handleChange(event) {
        setPicture(event.target.files[0])
    }

    const handleCheckboxChange = (id) => {
        const updatedCheckboxes = followers.map(checkbox =>
            checkbox.id == id ? { ...checkbox, checked: !checkbox.checked } : checkbox
        )
        setFollowers(updatedCheckboxes)
    }

    useEffect(() => {
        const fetchFollowers = async () => {
            try {
                const r = await followService.followersApi()
                for (let i in r.data) {
                    r.data[i].checked = false
                }
                setFollowers(r.data)
            } catch (error) {
                console.error("Error when fetching followers :", error.response.data.error)
                setErrors(error.response.data.error)
            }
        }
        fetchFollowers()
    }, [])
    
    return (
        <form className="post-form" onSubmit={handlePostSubmit}>
            <div className="form-div">
                <label className="form-label" htmlFor="post-title">Post title:</label>
                <input
                    id="post-title"
                    type="text"
                    value={postTitle}
                    onChange={(e) => setPostTitle(e.target.value)}
                />
            </div>
            <div className="form-div">
                <label className="form-label" htmlFor="post-content">Post message:</label>
                <input
                    id="post-content"
                    type="text"
                    value={postContent}
                    onChange={(e) => setPostContent(e.target.value)}
                />
            </div>
            { urlParams.split('/')[3] === "home" && (
                <div className="form-div">
                    <div>Privacy setting</div>
                    <label className="form-label" htmlFor="public">Public</label>
                    <input 
                        type="radio"
                        id="public"
                        name="postStatus"
                        value={0}
                        onChange={(e) => setPostPrivacy(e.target.value)}
                        required
                    />
                    <label className="form-label" htmlFor="semi-private">Semi Private</label>
                    <input
                        type="radio"
                        id="semi-private"
                        name="postStatus"
                        value={1}
                        onChange={(e) => setPostPrivacy(e.target.value)}
                    />
                    
                    <label className="form-label" htmlFor="private">Private</label>
                    <input
                        type="radio"
                        id="private"
                        name="postStatus"
                        value={2}
                        onChange={(e) => setPostPrivacy(e.target.value)}
                    />
                </div>
            )}
            { postPrivacy == 1 && (
                <div className="form-div">
                    <div>Select a follower than can see this post</div>
                    { followers.length > 0 ? (
                        followers.map((follower) => (
                            <div key={"follower_"+follower.id}>
                                <div>{follower.nickname}</div>
                                <input type="checkbox" onChange={() => handleCheckboxChange(follower.id)} />
                            </div>
                        ))
                    ) : (
                        <div className="comments">No followers.</div>
                    )}
                </div>
            )}
            <div className="form-div">
                <label className="form-label" htmlFor="picture">Picture: </label>
                <input
                    id="picture"
                    type="file"
                    onChange={handleChange}
                />
                {errors.picture && (
                    <small style="color: red">* {errors.picture}</small>
                )}
            </div>
            <button className="submit-form" type="submit">Submit</button>
        </form>
    )
}

export default PostForm