import PostList from '../../components/posts/post';
import { useParams } from 'react-router-dom';
import "../../css/post.css"

const UserPosts = () => {

    let userId = useParams().userId

    return (
        <div id="home-main-div">
            <PostList/>
        </div>
    )
}

export default UserPosts