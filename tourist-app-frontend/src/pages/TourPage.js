import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Container, Typography, Button } from '@material-ui/core';
import { getTour } from '../services/tourService';
import BookingForm from '../components/BookingForm';

const TourPage = () => {
    const { id } = useParams();
    const [tour, setTour] = useState(null);
    const [showBookingForm, setShowBookingForm] = useState(false);

    useEffect(() => {
        fetchTour();
    }, [id]);

    const fetchTour = async () => {
        try {
            const response = await getTour(id);
            setTour(response.data);
        } catch (error) {
            console.error('Error fetching tour: ', error);
        }
    };

    if (!tour) {
        return <Typography>Loading...</Typography>;
    }

    return (
        <Container>
            <Typography variant="h2" gutterBottom>
                {tour.name}
            </Typography>
            <Typography variant="body1" paragraph>
                {tour.description}
            </Typography>
            <Typography variant="h6" gutterBottom>
                Price: ${tour.price}
            </Typography>
            <Button
                variant="contained"
                color="primary"
                onClick={() => setShowBookingForm(!showBookingForm)}
            >
                {showBookingForm ? 'Hide Booking Form' : 'Book This Tour'}
            </Button>
            {showBookingForm && (
                <BookingForm
                    tourId={tour.id}
                    onBookingComplete={() => {
                        setShowBookingForm(false);
                        // Aqui podemos agregar uma notificacion de exito
                    }}
                />
            )}
        </Container>
    );
};

export default TourPage;
