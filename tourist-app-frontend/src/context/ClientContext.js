import React, { createContext, useContext, useState, useEffect } from 'react';
import { getClients } from '../services/clientService';

const ClientContext = createContext();

export const ClientProvider = ({ children }) => {
    const [getClients, setClients] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchClients();
    }, []);

    const fetchClients = async () => {
        try {
            const response = await getClients();
            setClients(response.data);
            setLoading(false);
        } catch (error) {
            console.error('Error fetching clients: ', error);
            setLoading(false);
        }
    };

    return (
        <ClientContext.Provider value={{ clients, loading, fetchClients }}>
            {children}
        </ClientContext.Provider>
    );
};

export const useClientContext = () => useContext(ClientContext);
