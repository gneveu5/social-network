import PostList from '../components/posts/post'
import EventList from '../components/group/eventList'
import GroupRegister from '../components/group/groupRegister'
import GroupMemberList from "../components/group/groupMemberList"
import { groupService } from "../_services/group.service"
import { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import "../css/post.css"

export const Group = () => {

    const [isMember, setIsMember] = useState(false)
    const [groupName, setGroupName] = useState("")
    const [groupId, setGroupId] = useState("")
    const navigate = useNavigate()

    useEffect(() => {

        const urlParams = window.location.href

        const fetchGroupExist = async () => {
            try {
                const response = await groupService.fetchGroupExist(urlParams.split('/')[4]) 
                if (response.data === "") {
                    navigate("/home")
                } else {
                    setGroupName(response.data["Title"])
                    setGroupId(response.data["Id"])
                }
            } catch (error) {
                console.error(error)
            }
        }

        const fetchIsGroupMember = async () => {
            try {
                const response = await groupService.fetchGroupMember(urlParams.split('/')[4])
                if (response.data) {
                    setIsMember(true)
                }
            } catch (error) {
                console.error(error)
            }
        }
        fetchGroupExist()
        fetchIsGroupMember()
    }, [])

    return (
        <div id="home-main-div">
            { isMember ? (
                <div>
                    <h1>{groupName} Group activity</h1>
                    <div>
                        <PostList/>
                    </div>
                    <div>
                        <EventList/>
                    </div>
                    <div>
                        <GroupMemberList groupId={groupId}/>
                    </div>
                </div>
            ) : (
                <div>
                    <h1>welcome to the group: {groupName}</h1>
                    <div>
                        <GroupRegister/>
                    </div>
                </div>
            )}
        </div> 
    )
}
  
export default Group