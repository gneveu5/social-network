import './css/App.css';
import {Routes, Route, Navigate} from "react-router-dom";
import { ProtectedRoute } from "./protectedroutes";
import ErrorPage from './pages/error';
// import { AuthProvider } from "./hooks/useAuth";


import { LoginPage } from "../src/pages/login";
import {Registerpage} from "./pages/register"

import MainLayout from "./pages/mainlayout";
import Home from "./pages/home";
import Group from "./pages/group";
import GroupCreation from "./pages/groupcreation";
import Notifications from "./pages/notifications"

import UserLayout from "./pages/user/layout";
import UserPosts from "./pages/user/posts";
import UserFollowers from "./pages/user/followers";
import UserFollowing from "./pages/user/following";
import UserGroups from "./pages/user/groups";

function App() {
  return (
      <div className='App'>
        <Routes>
          <Route path='/error' element={<ErrorPage />} />
          <Route path='/login' element={<LoginPage />} />
          <Route path='/register' element={<Registerpage />} />
          <Route path='*' element={<Navigate to="/home" />} /> 
          <Route path='/' element={<ProtectedRoute><MainLayout /></ProtectedRoute>}>
            <Route path='/' element={<Navigate to="/home" />} /> 
            <Route path="home" element={<ProtectedRoute><Home /></ProtectedRoute>}/>
            <Route path="notifications" element={<ProtectedRoute><Notifications /></ProtectedRoute>}/>
            <Route path="groupcreation" element={<ProtectedRoute><GroupCreation /></ProtectedRoute>} />
            <Route path="group/:groupId" element={<ProtectedRoute><Group /></ProtectedRoute>}/>
            <Route path="user/:userId/" element={<ProtectedRoute><UserLayout /></ProtectedRoute>}>
              <Route path='' element={<Navigate to="posts" />} /> 
              <Route path="posts" element={<ProtectedRoute><UserPosts /></ProtectedRoute>} />
              <Route path="followers" element={<ProtectedRoute><UserFollowers /></ProtectedRoute>} />
              <Route path="following" element={<ProtectedRoute><UserFollowing /></ProtectedRoute>} />
              <Route path="groups" element={<ProtectedRoute><UserGroups /></ProtectedRoute>} />
          </Route>
          </Route>
        </Routes>
      </div>
  );
}
export default App 
