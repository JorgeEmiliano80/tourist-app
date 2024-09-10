import React, { useState, useEffect } from 'react';
import { Container, Typography, Grid } from '@material-ui/core';
import TourCard from '../components/TourCard';
import { getTours } from '../services/tourService';

const HomePage = () => {
    const [tours, setTours] = useState([]);

    useEffect(() => {
        fetchTours();
    }, []);

    const fetchTours = async () => {
        try {
            const response = await getTours();
            setTours(response.data);
        } catch (error) {
            console.error('Error fetching tours: ', error);
        }
    };

    return (
        <Container>
            <Typography variant="h2" gutterBottom>
                Welcome to Tourist App
            </Typography>
            <Typography variant="h4" gutterBottom>
                Featured Tours
            </Typography>
            <Grid container spacing={3}>
                {tours.map((tour) => (
                    <Grid item xs={12} sm={6} md={4} key={tour.id}>
                        <TourCard tour={tour} />
                    </Grid>
                ))}
            </Grid>
        </Container>
    );
};

export default HomePage;
