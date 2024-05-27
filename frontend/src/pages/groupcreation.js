import { groupService } from "../_services/group.service"
import { useNavigate } from "react-router-dom";
import { useState } from "react"
import "../css/post.css"

export const GroupCreation = () => {

    const [groupTitle, setGroupTitle] = useState("")
    const [groupDescription, setGroupDescription] = useState("")
    const navigate = useNavigate();

    const handleGroupSubmit = async (e) => {

        e.preventDefault()

        let payload = {
            Title: groupTitle,
            Description: groupDescription,
        }
        await groupService.groupCreation(payload)
        .then((r) => {
            navigate("/group/" + r.data)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    return (
        <div id="home-main-div">
            <h1>Group creation</h1>
            <form className="group-form" onSubmit={handleGroupSubmit}>
                <div className="form-div">
                    <label className="form-label" htmlFor="group-title">Group title:</label>
                    <input
                        id="group-title"
                        type="text"
                        value={groupTitle}
                        onChange={(e) => setGroupTitle(e.target.value)}
                    />
                </div>
                <div className="form-div">
                    <label className="form-label" htmlFor="group-description">Group description:</label>
                    <input
                        id="group-description"
                        type="text"
                        value={groupDescription}
                        onChange={(e) => setGroupDescription(e.target.value)}
                    />
                </div>
                <button className="submit-form" type="submit">Submit</button>
            </form>
        </div> 
    )
}
  
export default GroupCreation