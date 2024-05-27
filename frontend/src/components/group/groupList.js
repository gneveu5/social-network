import { groupService } from "../../_services/group.service"
import { Link } from "react-router-dom"
import { useState, useEffect} from "react"
import "../../css/post.css"

export const GroupList = () => {

    const urlParams = window.location.href

    const [groups, setGroups] = useState([])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const r = await groupService.fetchGroupList(urlParams.split('/')[4])
                setGroups(r.data)
            } catch (e) {
                console.error("Error fetching data:", e)
            }
        }
        fetchData()
    }, [])

    return (
        <div>
            {groups && (groups.map((item) => (
                <div key={"event" + item.Id} id={item.Id} className="posts">
                    <div className="post-header">
                        <div className="header-subdiv">
                            {/* <a className="header-title" href={"/group/" + item.Id}>{item.Title}</a> */}
                            <Link to={"/group/" + item.Id} className="header-title">{item.Title}</Link>
                        </div>
                        <div className="header-div-name">{"group admin: " + item.AdminName}</div>
                    </div>
                    <div className="post-body">{item.Description}</div>
                </div>
            )))}
        </div>
    )
}
  
export default GroupList