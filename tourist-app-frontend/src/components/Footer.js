import React from 'react';
import { Typography, Container, Box } from '@mui/material';

const Footer = () => {
    return (
        <Box
            component="footer"
            sx={{
                marginTop: 'auto',
                py: 3, // padding top and bottom
                backgroundColor: 'background.paper',
            }}
        >
            <Container maxWidth="sm">
                <Typography variant="body1" align="center">
                    © {new Date().getFullYear()} Tourist App. All rights reserved.
                </Typography>
            </Container>
        </Box>
    );
};

export default Footer;