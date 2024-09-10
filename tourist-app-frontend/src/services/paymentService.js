import api from './api';

export const getPayments = () => api.get('/payments');
export const getPayment = (id) => api.get(`/payments/${id}`);
export const createPayment = (paymentData) => api.post('/payments', paymentData);
export const updatePayment = (id, paymentData) => api.put(`/payments/%{id}`, paymentData);
export const deletePayment = (id) => api.delete(`/payments/%{id}`);
