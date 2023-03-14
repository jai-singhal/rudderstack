import React, { useState } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';
import {createTrackingPlan} from "../../api/trackingplans"

const TrackingPlanCreateModal = (props) => {
const [name, setName] = useState('');
const [description, setDescription] = useState('');
const [eventRuleName, setEventRuleName] = useState('');
const [eventRuleDescription, setEventRuleDescription] = useState('');
const [eventRule, setEventRule] = useState('');

const handleSubmit = async (event) => {
	event.preventDefault();
	// Submit form data
	const data = {
		name: name,
		description: description,
		eventRule: {
			name: eventRuleName,
			description: eventRuleDescription,
			rules: JSON.parse(eventRule)
		}
	};
	console.log(data);
	try {
		const result = await createTrackingPlan(data);
		console.log(result);
		// Clear form fields
		setName('');
		setDescription('');
		setEventRuleName('');
		setEventRuleDescription('');
		setEventRule('');
		// Close modal
		props.onHide();
	} catch (error) {
		console.error(error);
	// Handle error
	}
}

return (
<Modal {...props} aria-labelledby="contained-modal-title-vcenter" centered>
	<Modal.Header closeButton>
		<Modal.Title id="contained-modal-title-vcenter">
			Add Tracking Plan
		</Modal.Title>
	</Modal.Header>
	<Modal.Body>
		<Form onSubmit={handleSubmit}>
			<Form.Group controlId="formName">
				<Form.Label>Name</Form.Label>
				<Form.Control type="text" value={name} onChange={(e)=> setName(e.target.value)}
					required
					/>
			</Form.Group>
			<Form.Group controlId="formDescription">
				<Form.Label>Description</Form.Label>
				<Form.Control as="textarea" rows={3} value={description} onChange={(e)=> setDescription(e.target.value)}
					required
					/>
			</Form.Group>
			<hr />
			<h5>Event Rule</h5>
			<Form.Group controlId="formEventRuleName">
				<Form.Label>Name</Form.Label>
				<Form.Control type="text" value={eventRuleName} onChange={(e)=> setEventRuleName(e.target.value)}
					required
					/>
			</Form.Group>
			<Form.Group controlId="formEventRuleDescription">
				<Form.Label>Description</Form.Label>
				<Form.Control as="textarea" rows={3} value={eventRuleDescription} onChange={(e)=>
					setEventRuleDescription(e.target.value)}
					required
					/>
			</Form.Group>
			<Form.Group controlId="formEventRule">
				<Form.Label>Rules</Form.Label>
				<Form.Control as="textarea" rows={3} value={eventRule} onChange={(e)=> setEventRule(e.target.value)}
					required
					/>
			</Form.Group>
			<Button type="submit">Submit</Button>
		</Form>
	</Modal.Body>
</Modal>
);
}

export default TrackingPlanCreateModal;
