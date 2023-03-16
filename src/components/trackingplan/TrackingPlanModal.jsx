import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import { createTrackingPlan, updateTrackingPlan } from '../../api/trackingplans';
import {getTrackingPlan} from "../../api/trackingplans"

const TrackingPlanModal = ({show, onHide, onSubmit, isUpdate = false, trackingPlanId = null}) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [initialData, setInitalData] = useState([{}, ]);
  const [eventsData, setEventsData] = useState([{
    name: '',
    description: '',
    rules: '',
  }, ]);
  const [error, setError] = useState('');

  useEffect(() => {
    if (show && isUpdate && trackingPlanId) {
		// let trackingplan = populateInitalData()
		getTrackingPlan(trackingPlanId).then((trackingplan) => {
			setInitalData(trackingplan)
			console.log("uodated", show, isUpdate, trackingplan)
			setName(trackingplan.name);
			setDescription(trackingplan.description);
			// console.log( "xxx")
			let eventsdata = []
			trackingplan.events.map((event ) => {
				eventsdata.push({
					name: event.name,
					description: event.description,
					rules: JSON.stringify(event.rules, null, 4)
				})
			})
			setEventsData(eventsdata);
		})

    //   setEventsData(initialData.rules.events);
    }
  }, [show, ]);

	// const  populateInitalData = async () =>  {
	// 	let trackingplan = await getTrackingPlan(trackingPlanId)
	// 	console.log(trackingplan)
	// 	return trackingplan
	// }

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
        let result;
        if (isUpdate) {
          result = await updateTrackingPlan(initialData.id, data);
        } else {
          result = await createTrackingPlan(data);
        }
        const { error } = result;
        if (error) {
          setError(error.error);
          return false;
        }
        onHide();
        onSubmit(data);
        setName('');
        setError(null);
        setDescription('');
        setEventsData([{
          name: '',
          description: '',
          rules: '',
        }, ]);
      } catch (error) {
        console.error(error);
        setError('An error occurred ' + error);
        return false;
      }
    }
  };

  	return (
    <Modal show={show} onHide={onHide} aria-labelledby="tracking-plan-create-modal" size="lg">
    	<Modal.Header closeButton>
    		<Modal.Title id="tracking-plan-create-modal">
			{isUpdate ? 'Update Tracking Plan' : 'Add Tracking Plan'}
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
    				<Form.Control as="textarea" rows={1} value={description} onChange={(e)=>
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
    					<Form.Control as="textarea" rows={1} value={eventRule.description} onChange={(e)=> {
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
    					<Form.Control as="textarea" rows={6} value={eventRule.rules} onChange={(e)=> {
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
    			<Button type="submit" className="float-end">{isUpdate ? 'Update' : 'Add'}</Button>
    		</Form>
    	</Modal.Body>
    </Modal>
	);
}



export default TrackingPlanModal;