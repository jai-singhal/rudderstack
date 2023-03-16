import { Offcanvas, Card, Button } from 'react-bootstrap';
import { useState } from 'react';
import EventCreateModal from './EventModal';

function EventRulesCard({ rule }) {
	const [showRules, setShowRules] = useState(false);
	const [showCreateEventModal, setShowCreateEventModal] = useState(false);

	const toggleShowRules = () => {
		setShowRules(!showRules);
	};

	return (
		<Card>
			<Card.Body>
				<Card.Title>{rule.name}</Card.Title>
				<Card.Subtitle className="mb-2 text-muted">Created at: {rule.created_at}</Card.Subtitle>
				<Card.Text>
					Description: {rule.description}<br />
					<Button variant="button" onClick={toggleShowRules}>
						{showRules ? 'Hide Rules' : 'Show Rules'}
					</Button>
				</Card.Text>
				{
				showRules &&
				<pre>{JSON.stringify(rule.rules, null, 2)}</pre>
				}
			</Card.Body>
			<Card.Footer>
				<Button variant="primary" onClick={setShowCreateEventModal}>
					Add Event
				</Button>
				<EventCreateModal show={showCreateEventModal} onHide={()=> setShowCreateEventModal(false)}
					onSubmit={(data) => console.log(data)}
					eventName = {rule.name}
					/>
			</Card.Footer>
		</Card>
	);
}

function EventRulesOffCanvas({show, onHide, trackingPlanDetail}) {
	const handleClose = () => {
		onHide();
	};

	return (
		<Offcanvas show={show} onHide={handleClose} placement="end" style={{"width":"40%"}}>
			<Offcanvas.Header closeButton>
				<Offcanvas.Title>Event Rules</Offcanvas.Title>
			</Offcanvas.Header>
			<Offcanvas.Body>
				{trackingPlanDetail.map((rule) => (
				<div className="col-12" key={rule.id}>
					<EventRulesCard rule={rule} />
				</div>
				))}
			</Offcanvas.Body>
		</Offcanvas>
	);
}

export default EventRulesOffCanvas;