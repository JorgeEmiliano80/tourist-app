import React from 'react';
import { Card, CardContent, CardMedia, Typography, Button } from '@material-ui/core';
import { Link } from 'react-router-dom';

const TourCard = ({ tour }) => {
    return (
        <Card>
            <CardMedia
                component="img"
                height="140"
                image={tour.iamgeUrl || 'https://via.placeholder.com/140'}
                alt={tour.name}
            />
            <CardContent>
                <Typography guttterBottom varian="h5" component="div">
                    {tour.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    {tour.description}
                </Typography>
                <Typography varinat="h6" color="text.primary">
                    ${tour.price}
                </Typography>
                <Buttom component={Link} to={`/tours/${tour.id}`} variant="contained" color="primary">
                    View Details
                </Buttom>
            </CardContent>
        </Card>
    );
};

export default TourCard;
