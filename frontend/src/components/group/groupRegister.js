import { groupService } from "../../_services/group.service"
import { useState, useEffect} from "react"
import "../../css/post.css"

export const GroupRegister = (handleGroupAskMembershipReturn) => {

    const urlParams = window.location.href
    const [hasAsked, setHasAsked] = useState(false);

    const handleGroupAskMembership = async () => {

        let payload
        payload = {
            groupId: urlParams.split('/')[4],
        }

        await groupService.groupAskMembership(payload)
        .then((r) => {
            setHasAsked(r.data)
        })
        .catch((e) => {
            if (e.response.data !== undefined) {
                alert(e.response.data.error)
            } else {
                alert("problem server side")
            }
        })
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const r = await groupService.fetchGroupAskMembership(urlParams.split('/')[4])
                setHasAsked(r.data)
            } catch (e) {
                console.error("Error fetching data:", e)
            }
        }
        fetchData()
    }, [])

    return (
        <div>
            {!hasAsked ? (
                <button className="group-Register" onClick={() => handleGroupAskMembership()}>Ask to join group</button>
            ) : (
                <div>Your request is being reviewed</div>
            )}
        </div>
    )
}
  
export default GroupRegister