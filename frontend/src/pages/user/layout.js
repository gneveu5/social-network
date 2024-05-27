import { Link, useParams, Outlet, useLocation } from 'react-router-dom';
import React, { useState } from 'react';

import UserProfile from './profile.js'

//style
import '../../css/UserLayout.css'; 

const UserLayout = () => {
    const userId = useParams().userId
    const location = useLocation();
    let a = useParams()

    const [isProfilePublic, setIsProfilePublic] = useState(false); 
    
    const adr = "/user/" + userId + "/"

    const handleProfileInfo = (isPublic, followStatus) => {
        if (followStatus == 1 || followStatus == 3) {
            setIsProfilePublic(false)

        }else {
            setIsProfilePublic(isPublic);


        }
    };




    return (
        <div className="user-layout-container">
            <UserProfile userId={userId} onProfileInfo={handleProfileInfo} />            <div>
            {!isProfilePublic && ( 
            <nav>
            <ul className='usernavbar'>
                <li className='usernavbarchild'>
                    <Link to={adr + "posts"} className={location.pathname === adr + "posts" ? 'active' : ''}>Posts</Link>
                </li>
                <li className='usernavbarchild'>
                    <Link to={adr + "followers"} className={location.pathname === adr + "followers" ? 'active' : ''}>Followers</Link>
                </li>
                <li className='usernavbarchild'>
                    <Link to={adr + "following"} className={location.pathname === adr + "following" ? 'active' : ''}>Following</Link>
                </li>
                <li className='usernavbarchild'>
                    <Link to={adr + "groups"} className={location.pathname === adr + "groups" ? 'active' : ''}>Groups</Link>
                </li>
            </ul>
        </nav>
            )}
            </div>
            <Outlet />
        </div>

    )
}

export default UserLayout