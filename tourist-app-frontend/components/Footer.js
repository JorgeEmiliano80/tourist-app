import React from 'react';
import { Typography, Container } from '@material-ui/core';

const Footer = () => {
    return (
        <footer style={{ marginTop: 'auto', padding: '20px 0', backgroundColor: '#f5f5f5' }}>
            <Container maxWidth="sm">
                <Typography variant="body1" align="center">
                    © {new Date().getFullYear()} Tourist App. All rights resrved.
                </Typography>
            </Container>
        </footer>
    );
};

export default Footer;
