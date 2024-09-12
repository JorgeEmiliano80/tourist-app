import React from 'react';
import { Card, CardContent, CardMedia, Typography, Button, Box } from '@mui/material';
import { Link } from 'react-router-dom';

const TourCard = ({ tour }) => {
    return (
        <Card sx={{ maxWidth: 345, height: '100%', display: 'flex', flexDirection: 'column' }}>
            <CardMedia
                component="img"
                height="140"
                image={tour.imageUrl || 'https://via.placeholder.com/140'}
                alt={tour.name}
            />
            <CardContent sx={{ flexGrow: 1 }}>
                <Typography gutterBottom variant="h5" component="div">
                    {tour.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    {tour.description}
                </Typography>
                <Typography variant="h6" color="text.primary" sx={{ mt: 2 }}>
                    ${tour.price}
                </Typography>
            </CardContent>
            <Box sx={{ p: 2 }}>
                <Button 
                    component={Link} 
                    to={`/tours/${tour.id}`} 
                    variant="contained" 
                    color="primary"
                    fullWidth
                >
                    View Details
                </Button>
            </Box>
        </Card>
    );
};

export default TourCard;