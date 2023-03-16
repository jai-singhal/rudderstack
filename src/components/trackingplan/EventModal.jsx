import { useState } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import { createEvent } from '../../api/events';

function EventCreateModal({
    show,
    onHide,
    onSubmit,
    eventName
}) {
    const [name, setName] = useState(eventName || '');
    const [properties, setProperties] = useState('');
    const [error, setError] = useState(null);
    

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            JSON.parse(properties)
        } catch (error) {
            setError('Event properties is not valid JSON. Please correct.');
            return false;
        }
        const data = {
            name: name,
            properties: JSON.parse(properties),
        };

        try {
            const {error, _} = await createEvent(data);
            if (error) {
                setError(error.error);
                return false;
            }
            if(eventName.length === 0){
                setName('')
            }
            setProperties('')
            setError(null)
            onHide();
            onSubmit(data);
        } catch (error) {
            console.error(error);
            setError('An error occurred ' + error);
            return false;
        } 
    };

  return (
    <Modal show={show} onHide={onHide} aria-labelledby="event-create-modal">
        <Modal.Header closeButton>
            <Modal.Title>Add Event</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            {error && <Alert variant="danger">{error}</Alert>}
            <Form onSubmit={handleSubmit}>
                <Form.Group controlId="formName">
                    <Form.Label>Name</Form.Label>
                    <Form.Control type="text" placeholder="Enter name" value={eventName?eventName:name}
                        onChange={(event)=> setName(event.target.value)}
                        disabled={eventName.length > 0}
                        />
                </Form.Group>

                <Form.Group controlId="formProperties">
                    <Form.Label>Properties</Form.Label>
                    <Form.Control as="textarea" placeholder="Enter JSON properties" value={properties}
                        onChange={(event)=> setProperties(event.target.value)}
                        />
                </Form.Group>
            </Form>
        </Modal.Body>
        <Modal.Footer>
            <Button variant="secondary" onClick={onHide}>
                Close
            </Button>
            <Button variant="primary" onClick={handleSubmit}>
                Add
            </Button>
        </Modal.Footer>
    </Modal>
  );
}

export default EventCreateModal;
