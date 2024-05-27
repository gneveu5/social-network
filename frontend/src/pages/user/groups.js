import GroupList from '../../components/group/groupList';
import { useParams } from 'react-router-dom';

const UserGroups = () => {
    let userId = useParams().userId
    return (
        <div>
            <GroupList/>
        </div>
    )
}

export default UserGroups