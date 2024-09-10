import api from './api';

export const getTours = () => api.get('/tours');
export const getTour = (id) => api.get(`/tours/${id}`);
export const createTour = (tourData) => api.post('/tours', tourData);
export const updateTour = (id, tourData) => api.put(`/tours/${id}`, tourData);
export const deleteTour = (id) => api.delete(`/tours/${id}`);
