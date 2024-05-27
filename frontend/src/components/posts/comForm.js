import { postService } from "../../_services/post.service"
import { useState} from "react"
import "../../css/post.css"

export const ComForm = ({itemId, handleComSubmitReturn}) => {

    const [commentInputValue, setCommentInputValues] = useState({})
    const [picture, setPicture] = useState("")
    const [errors, setErrors] = useState({})

    const handleCommentSubmit = async (e, id) => {
        e.preventDefault()
        let payload = {
            postId: id,
            message: commentInputValue[id+"comment-content"],
            picture: picture,
        }
        console.log(payload)
        await postService.commenting(payload)
        .then((r) => {
            handleComSubmitReturn(r.data)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    const handleChange = (event, id) => {
        setCommentInputValues({ ...commentInputValue, [id]: event.target.value })
    }

    function handlePictureChange(event) {
        setPicture(event.target.files[0])
    }
    
    return (
        <form className="comment-form" onSubmit={(event) => handleCommentSubmit(event, itemId)}>
            <div>
                <label htmlFor={itemId + "comment-content"}>Reply to this post:</label>
                <input
                    id={itemId + "comment-content"}
                    type="text"
                    value={commentInputValue[itemId + "comment-content"] || ""}
                    onChange={(event) => handleChange(event, itemId + "comment-content")}
                    required
                />
            </div>
            <div className="form-div">
                <label className="form-label" htmlFor="picture">Picture: </label>
                <input
                    id="picture"
                    type="file"
                    onChange={handlePictureChange}
                />
                {errors.picture && (
                    <small style="color: red">* {errors.picture}</small>
                )}
            </div>
            <button type="submit">Submit</button>
        </form>
    )
}
  
export default ComForm
