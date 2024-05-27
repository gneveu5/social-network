import { postService } from "../../_services/post.service"
import { useState, useEffect } from "react"
import PostForm from './postForm'
import ComForm from './comForm'
import MyImageComponent from '../profilepicture';
import "../../css/post.css"

export const PostList = () => {

    const urlParams = window.location.href

    const [posts, setPosts] = useState([])
    const [comments, setComments] = useState([])

    const handleCommentFetch = async (e) => {

        const commentExists = comments.some((comment) => comment.id === e.Id)
        if (!commentExists) {
            e.Opened = "Close comments"
            let payload = {
                postId: e.Id
            }
            await postService.postCommentFetch(payload)
            .then((r) => {
                setComments(comments => [...comments, { id: e.Id, data: r.data }])
            })
            .catch((err) => {
                if (err.response.data !== undefined) {
                    alert(err.response.data.error)
                } else {
                    alert("problem server side")
                }
            })
        } else {
            e.Opened = "See comments"
            setComments(comments.filter((comment) => comment.id !== e.Id));
        }
    }

    const handlePostSubmitReturn = async (e) => {
        setPosts(e)
    }

    const handleComSubmitReturn = async (e) => {
        setComments(comments.map((comment) => {
            if (comment.id === e[0].PostId) {
                return { ...comment, data: e }
            }
            return comment
        }))
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                let tmp
                if (urlParams.split('/')[3] === "home") {
                    const r = await postService.postfetch()
                    tmp = r.data
                } else if (urlParams.split('/')[3] === "group") {
                    const r = await postService.postfetch(urlParams.split('/')[4])
                    tmp = r.data
                } else {
                    const r = await postService.postfetch("profil" + urlParams.split('/')[4])
                    tmp = r.data
                }
                for (let i in tmp) {
                    tmp[i]["Opened"] = "See comments"
                }
                setPosts(tmp)
            } catch (e) {
                console.error("Error fetching data:", e)
            }
        }
        fetchData()
    }, [])
    
    return (
        <div className="post-form-list">
            { (urlParams.split('/')[3] === "home" || urlParams.split('/')[3] === "group") && (
                <PostForm handlePostSubmitReturn={handlePostSubmitReturn}/>
            )}
            {posts && (posts.map((item) => (
                <div key={"post" + item.Id} id={item.Id} className="posts">
                    <div className="post-header">
                        <div className="header-subdiv">
                            <div className="header-title">{item.Title}</div>
                            <div className="header-div-name">{"by " + item.AuthorName}</div>
                        </div>
                        <div className="header-date">{item.CreatedAt.replace(/[a-z]/gi, ' ')}</div>
                    </div>
                    <div className="post-content">
                        <div className="post-body">{item.Body}</div>
                        { item.Img.String && (
                            <div className="post-img">
                                <MyImageComponent avatarPath={item.Img["String"]} avatarSize={150} route={"img"}/>
                            </div>
                        )}
                    </div>
                    <button className="see-comments" onClick={() => handleCommentFetch(item)}>{item.Opened}</button>
                    {comments && (comments.map((comment) => comment.id === item.Id && (
                        <div key={"post-comments" + item.Id}>
                            { (urlParams.split('/')[3] === "home" || urlParams.split('/')[3] === "group") && (
                                <ComForm itemId={item.Id} handleComSubmitReturn={handleComSubmitReturn}/>
                            )}
                            {comment.data && comment.data.length > 0 ? (
                                comment.data.map((com) => ( 
                                    <div key={"com" + com.Id} className="comments">
                                        <div className="com-header">
                                            <div className="header-div-name">{"by " + com.UserName}</div>
                                            <div className="header-date">{com.CreatedAt.replace(/[a-z]/gi, ' ')}</div>
                                        </div>
                                        <div className="com-content">
                                            <div>{com.Body}</div>
                                            { com.Img["String"] && (
                                                <div className="post-img">
                                                    <MyImageComponent avatarPath={com.Img["String"]} avatarSize={150} route={"img"}/>
                                                </div>
                                            )}
                                        </div>
                                    </div>
                                ))
                            ) : (
                                <div className="comments">No comment for this post.</div>
                            )}
                        </div>
                    )))}
                </div>
            )))}
        </div>
    )
}
  
export default PostList