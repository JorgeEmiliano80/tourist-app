import api from './api';
import { updateTour } from './tourService';

export const getBookings = () => api.get('./bookings');
export const getBooking = (id) => api.get(`/bookings/${id}`);
export const createBooking = (bookingData) => api.post('/bookings', bookingData);
export const updateBooking = (id, bookingData) => api.put(`/bookings/${id}`, bookingData);
export const deleteBooking = (id) => api.delete(`/bookings/${id}`);
