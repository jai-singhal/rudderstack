import React, { useState } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import { createTrackingPlan } from '../../api/trackingplans';

const TrackingPlanCreateModal = ({show, onHide, onSubmit}) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [eventsData, setEventsData] = useState([{
  	name: '',
  	description: '',
  	rules: '',
  }, ]);
  const [error, setError] = useState('');

  const handleAddEvent = () => {
  	const newEvent = [...eventsData];
  	newEvent.push({
  		name: '',
  		description: '',
  		rules: '',
  	});
  	setEventsData(newEvent);
  };

  const handleRemoveEventRule = () => {
  	if (eventsData.length <= 1) return false;
  	const newEvent = [...eventsData];
  	newEvent.splice(eventsData.length - 1, 1);
  	setEventsData(newEvent);
  };

  const handleSubmit = async (e) => {
  	e.preventDefault();
  	let hasError = false;
  	const newEventsData = [...eventsData];

  	newEventsData.forEach((event, i) => {
  		try {
  			JSON.parse(event.rules);
  		} catch (error) {
  			setError(`Event ${i + 1}: '${event.name}' rule is not valid json. Please correct.`);
  			hasError = true;
  		}
  	});

  	if (!hasError) {
  		const data = {
  			tracking_plan: {
  				display_name: name,
  				description: description,
  				rules: {
  					events: eventsData.map((event) => ({
  						name: event.name,
  						description: event.description,
  						rules: JSON.parse(event.rules),
  					})),
  				},
  			},
  		};

  		try {
			const {error, result}  = await createTrackingPlan(data);
			if (error) {
                setError(error.error);
                return false;
            }
			onHide();
			onSubmit(data);
			setName('');
            setError(null)
			setDescription('');
			setEventsData([{
				name: '',
				description: '',
				rules: '',
			}, ]);
  		} catch (error) {
			console.error(error);
            setError('An error occurred '+ error);
            return false;
  		}
  	}
  };

  	return (
    <Modal show={show} onHide={onHide} aria-labelledby="tracking-plan-create-modal" size="lg">
    	<Modal.Header closeButton>
    		<Modal.Title id="tracking-plan-create-modal">
    			Add Tracking Plan
    		</Modal.Title>
    	</Modal.Header>
    	<Modal.Body>
			{error && <Alert variant="danger">{error}</Alert>}

    		<Form onSubmit={handleSubmit}>
    			<Form.Group controlId="formName">
    				<Form.Label>Name</Form.Label>
    				<Form.Control type="text" value={name} onChange={(e)=> setName(e.target.value)}
    					required
    					/>
    			</Form.Group>
    			<Form.Group controlId="formDescription">
    				<Form.Label>Description</Form.Label>
    				<Form.Control as="textarea" rows={3} value={description} onChange={(e)=>
    					setDescription(e.target.value)}
    					required
    					/>
    			</Form.Group>
    			<hr />
    			<br />
				<h5>Events</h5>
    			{eventsData.map( (eventRule, index) => (
    			<div key={index}>
    				<Form.Group controlId={`formEventRuleName-${index}`}>
    					<Form.Label>Name</Form.Label>
    					<Form.Control type="text" value={eventRule.name} onChange={(e)=> {
    						const newEvent = [...eventsData];
    						newEvent[index].name = e.target.value;
    						setEventsData(newEvent);
    						}}
    						required
    						/>
    				</Form.Group>
    				<Form.Group controlId={`formEventRuleDescription-${index}`}>
    					<Form.Label>Description</Form.Label>
    					<Form.Control as="textarea" rows={2} value={eventRule.description} onChange={(e)=> {
    						const newEvent = [...eventsData];
    						newEvent[index].description = e.target.value;
    						setEventsData(newEvent);
    						}}
    						required
							placeholder="Enter Event rule here..."
    						/>
    				</Form.Group>
    				<Form.Group controlId={`formEventRuleRules-${index}`}>
    					<Form.Label>Rules</Form.Label>
    					<Form.Control as="textarea" rows={4} value={eventRule.rules} onChange={(e)=> {
								const newEvent = [...eventsData];
								newEvent[index].rules = e.target.value;
								setEventsData(newEvent);
    						}}
    						required
							placeholder="Enter Event rule here..."
    						/>
    				</Form.Group>
					<br/>
    			</div>
    			))}
				<Button variant="primary" onClick={()=>handleAddEvent()}>+</Button>
				<Button variant="danger" disabled={eventsData.length <= 1} onClick={()=>handleRemoveEventRule()}>-</Button>

				<br/>
				<br/>
    			<Button type="submit" className="float-end">Submit</Button>
    		</Form>
    	</Modal.Body>
    </Modal>
	);
}


export default TrackingPlanCreateModal;