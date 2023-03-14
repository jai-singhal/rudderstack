import React, { useState, useEffect } from 'react';
import { Table, Button } from 'react-bootstrap';
import { getAllTrackingPlans, getTrackingPlan } from "../../api/trackingplans"
import EventRulesModal from './EventRulesModal';

const TrackingPlanTable = () => {
    const [trackingPlans, setTrackingPlans] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [trackingPlanDetail, setTrackingPlanDetail] = useState([]);
  
    const handleShowEventRules = async (id) => {
      if (trackingPlanDetail.length > 0 && trackingPlanDetail[0].tracking_plan_id === id) {
        return;
      }
      const eventRules = await getTrackingPlan(id);
      setTrackingPlanDetail(eventRules.rules)
      setShowModal(true);
    };

    useEffect(() => {
        fetchTrackingPlans();
    }, []);

    const fetchTrackingPlans = async () => {
        let trackingplans = await getAllTrackingPlans();
        setTrackingPlans(trackingplans.items)
    };

    return (
        <Table>
            <thead>
                <tr>
                    <th>S no.</th>
                    <th>Display Name</th>
                    <th>Description</th>
                    <th>Event Rules</th>
                </tr>
            </thead>
            <tbody>
                {trackingPlans.map((plan) => (
                <tr key={plan.id}>
                    <td>{plan.id}</td>
                    <td>{plan.display_name}</td>
                    <td>{plan.description}</td>
                    <td>
                    {/* <Button onClick={handleShowEventRules}>Click me!</Button> */}
                    <Button onClick={() => handleShowEventRules(plan.id)}>Click me!</Button>
                    <EventRulesModal
                        showModal={showModal}
                        onHide={() => setShowModal(false)}
                        trackingPlanDetail={trackingPlanDetail}
                    />
                    {/* <Button onClick={(e)=> {e.preventDefault(); handleShowEventRules(plan.id)}}>Click me!</Button> */}
                    {/* <Accordion>
                        <Accordion.Collapse>
                            <Card>
                                <Card.Body>
                                    <Table>
                                        <thead>
                                            <tr>
                                                <th>ID</th>
                                                <th>Name</th>
                                                <th>Description</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {trackingPlanDetail.map((rule) => (
                                            <tr key={rule.id}>
                                                <td>{rule.id}</td>
                                                <td>{rule.name}</td>
                                                <td>{rule.description}</td>
                                            </tr>
                                            ))}
                                        </tbody>
                                    </Table>
                                </Card.Body>
                            </Card>
                        </Accordion.Collapse>
                    </Accordion> */}
                    </td>
                </tr>
                ))}
            </tbody>
        </Table>
    );
};

export default TrackingPlanTable;
