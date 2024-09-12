import React, { useState, useEffect } from 'react';
import { Container, Typography } from '@mui/material';
import Grid2 from '@mui/material/Unstable_Grid2'; // Importamos Grid2
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
            <Typography variant="h2" component="h1" gutterBottom>
                Welcome to Tourist App
            </Typography>
            <Typography variant="h4" component="h2" gutterBottom>
                Featured Tours
            </Typography>
            <Grid2 container spacing={3}>
                {tours.map((tour) => (
                    <Grid2 xs={12} sm={6} md={4} key={tour.id}>
                        <TourCard tour={tour} />
                    </Grid2>
                ))}
            </Grid2>
        </Container>
    );
};

export default HomePage;