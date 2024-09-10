import React from 'react';
import { Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Button } from '@material-ui/core';
import { useAuthContext } from '../context/AuthContext';

const Header = () => {
    const { user, logout } = useAuthContext();

    return (
        <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" style={{ flexGrow: 1 }}>
                    Tourist App
                </Typography>
                <Button color="inherit" component={Link} to="/">Home</Button>
                <Button color="inherit" component={Link} to="/tours">Tours</Button>
                <Button color="inherit" component={Link} to="/bookings">Bookings</Button>
                <Button color="inherit" component={Link} to="/clients">Clients</Button>
                {user ? (
                    <Button color="inherit" onClick={logout}>Logout</Button>
                ) : (
                    <Button color="inherit" component={Link} to="/login">Login</Button>
                )}
            </Toolbar>
        </AppBar>
    );
};

export default Header;
