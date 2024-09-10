import React from 'react';
import { List, ListItem, ListItemText, ListItemSecondaryAction, IconButton } from '@material-ui/core';
import { Delete, Edit } from '@material-ui/icons';
import { useClientContext } from '../context/ClientContext';

const ClientList = ({ onEdit }) => {
    const { clients, fetchClients } = useClientContext();

    const handleDelete = async (id) => {
        try {
            await deleteClient(id);
            fetchClients();
        } catch (error) {
            console.error('Error deleting client: ', error);
        }
    };

    return (
        <List>
            {clients.map((client) => (
                <ListItem key={client.id}>
                    <ListItemText primary={client.name} secondary={client.email} />
                    <ListItemSecondaryAction>
                        <IconButton edge="end" aria-label="edit" onClick={() => onEdit(client)}>
                            <Edit />
                        </IconButton>
                        <IconButton edge="end" aria-label="delete" onClick={() => handleDelete(client.id)}>
                            <Delete />
                        </IconButton>
                    </ListItemSecondaryAction>
                </ListItem>
            ))}
        </List>
    );
};

export default ClientList;
