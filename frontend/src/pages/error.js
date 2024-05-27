import React from 'react';
import { useLocation, Link } from 'react-router-dom';

function ErrorPage() {
    const location = useLocation();
    const errorMessage = new URLSearchParams(location.search).get('message');

    return (
        <div>
            <h1>Error Page</h1>
            <p>{errorMessage}</p>
            <Link to="/home">Home</Link>

            </div>
            
    );
}

export default ErrorPage;
