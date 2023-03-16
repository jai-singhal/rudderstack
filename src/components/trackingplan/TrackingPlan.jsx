import React, { useState } from 'react';
import Header from '../Header'
import TrackingPlanTable from './TrackingPlanTable';
import TrackingPlanModal from './TrackingPlanModal'
import EventCreateModal from './EventModal'
import { Container, Button } from 'react-bootstrap';

const TrackingPlan = () => {
  const pageTitle = 'Tracking Plan'
  const pageDescription = 'Create your new Tracking Plan'
  const [showModal, setShowModal] = useState(false);
  const [showCreateEventModal, setShowCreateEventModal] = useState(false);
  const [tableKey, setTableKey] = useState(0);
  
  const refreshTable = () => {
    setTableKey(prevKey => prevKey + 1);
  }
  
  return (
    <div>
      <Header head={pageTitle} description={pageDescription} />
        <Container>
          <Button onClick={() => setShowModal(true)} className="float-end">Add Tracking Plan</Button>
          <TrackingPlanModal
              show={showModal}
              onHide={() => setShowModal(false)}
              onSubmit={(data) => {console.log(data);  refreshTable();}}
          />
          <Button variant="primary" onClick={setShowCreateEventModal} className="float-end">
                Add Event
            </Button>
            <EventCreateModal show={showCreateEventModal} onHide={()=> setShowCreateEventModal(false)}
                onSubmit={(data) => {console.log(data);}}
                eventName = {""}
            />
          <TrackingPlanTable key={tableKey} />
        </Container>

    </div>
  )
}

export default TrackingPlan