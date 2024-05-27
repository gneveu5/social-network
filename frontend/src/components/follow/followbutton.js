import React, { useState, useEffect } from 'react';
import { followService } from '../../_services/follow';
import { accountService } from '../../_services/account.service';
import '../../css/UserLayout.css';


const FollowButton = ({ followStatus, isprivate, nickname }) => {
    const [buttonText, setButtonText] = useState('');
  
    useEffect(() => {
      switch (followStatus) {
        case 0:
          setButtonText('Annuler demande');
          break;
        case 1:
          setButtonText('Unfollow');
          break;
        case 2:
          setButtonText(isprivate ? 'Envoyer une demande' : 'Follow');
          break;
        case 3:
           setButtonText(isprivate ? 'Passer le compte en public' : 'Passer le compte en privé');
          break;
        default:
          setButtonText('Erreur');
          break;
      }
    }, [followStatus, isprivate]);
  
    const handleButtonClick = () => {
      switch (buttonText) {
        case 'Envoyer une demande':
            followService.reqfollowprivate(nickname)
            .then(() => {setButtonText('Annuler demande')})
            .catch((error) => {console.log(error)});
          break;
        case 'Unfollow':
            followService.requnfollow(nickname)
            .then(() => {setButtonText(isprivate ? 'Envoyer une demande' : 'Follow')})
            .catch((error) => {console.log(error)});
          break;
        case 'Follow':
          followService.reqfollow(nickname)
            .then(() => {setButtonText('Unfollow');})
            .catch((error) => {console.log(error)});
          break;
        case 'Annuler demande':
            followService.reqcancelfollowprivate(nickname)
            .then(() => {setButtonText('Envoyer une demande');})
            .catch((error) => {console.log(error)});
          break;
        case 'Passer le compte en public':
            accountService.privatetopublic()
            .then(() => {setButtonText('Passer le compte en privé');})
            .catch((error) => {console.log(error)});
          break;
        case 'Passer le compte en privé':
            accountService.publictoprivate()
            .then(() => {setButtonText('Passer le compte en public');})
            .catch((error) => {console.log(error)});
          break;
        default:
          setButtonText('Erreur');
          break;
      }
    };
  
    return (
      <button className='butonn-profile' onClick={handleButtonClick}>
        {buttonText}
      </button>
    );
  };
  
  export default FollowButton;