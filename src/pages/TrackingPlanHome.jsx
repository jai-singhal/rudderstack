import React, { useState } from 'react';
import Header from '../components/Header'
import TrackingPlanTable from '../components/trackingplan/TrackingPlanTable';
import TrackingPlanCreateModal from '../components/trackingplan/TrackingPlanModal'
import { Container } from 'react-bootstrap';

const TrackingPlanHome = () => {
  // page content
  const pageTitle = 'Tracking Plan'
  const pageDescription = 'Create your new Tracking Plan'
  // eslint-disable-next-line no-undef
  const [showModal, setShowModal] = useState(false);

  return (
    <div>
      <Header head={pageTitle} description={pageDescription} />

        <Container>
          <button onClick={() => setShowModal(true)} className="float-right">Add Tracking Plan</button>

          <TrackingPlanTable>
          </TrackingPlanTable>
          
          <TrackingPlanCreateModal
              show={showModal}
              onHide={() => setShowModal(false)}
              onSubmit={(data) => console.log(data)}
          />
        </Container>

    </div>
  )
}

export default TrackingPlanHome