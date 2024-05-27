function MyImageComponent({ avatarPath, avatarSize, route }) {

    if (route == "avatar") {
        return (
            <>
                {avatarPath == "" || avatarPath == undefined ?( 
                    <img
                        src="https://cinbiose.uqam.ca/wp-content/uploads/sites/24/blank-profile-picture-g91ef0370b_1280.png"
                        alt="profil picture"
                        style={{ borderRadius: '50%', width: `${avatarSize}px`, height: `${avatarSize}px` }} 
                    />
                ):( 
                    <img
                        src={`http://localhost:8080/${route}/${avatarPath}`}
                        alt="profil picture"
                        style={{ borderRadius: '50%', width: `${avatarSize}px`, height: `${avatarSize}px` }} 
                    />
                )}
            </>
        )
    } else {
        return (
            <img
                src={`http://localhost:8080/${route}/${avatarPath}`}
                alt="profil picture"
                style={{ width: `${avatarSize}px`, height: `${avatarSize}px` }} 
            />
        )
    }
}
  
export default MyImageComponent
