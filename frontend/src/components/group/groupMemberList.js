import { useState, useEffect} from "react"
import { followService } from "../../_services/follow"
import { groupService } from "../../_services/group.service"
import "../../css/post.css"

export const GroupMemberList = (groupId) => {

    // const urlParams = window.location.href

    const [response, setResponse] = useState([])
    const [error, setError] = useState(null)

    const handleGroupInvitation = async (e) => {

        let payload = {
            Target: e,
            GroupId: groupId["groupId"],
        }
        await groupService.groupInviteMembership(payload)
        .then(() => {
            const tmp = response.filter(element => element.id !== e)
            setResponse(tmp)
        })
        .catch((err) => {
            if (err.response.data !== undefined) {
                alert(err.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    useEffect(() => {
        const fetchFollowers = async () => {
            try {
                const res = await followService.followersNotInGroup(groupId["groupId"])
                setResponse(res.data)
            } catch (error) {
                console.error("Error when fetching followers :", error.response.data.error)
                setError(error.response.data.error)
            }
        }

        fetchFollowers()
    }, [])

    return (
        <>
            {response && response.length > 0 ? (
                <div>
                    <div>Invite to the group</div>
                    <ul className='listeusers'>
                        {response.map(follower => (
                            <li key={follower.id}>
                                <button className="group-Register" onClick={() => handleGroupInvitation(follower.id)}>{follower.nickname}</button>
                            </li>
                        ))}
                    </ul>
                </div>
            ) : (
                <p>{error}</p>
            )}
        </>
    )
}
  
export default GroupMemberList
