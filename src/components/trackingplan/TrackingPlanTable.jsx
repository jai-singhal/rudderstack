import React, { useState, useEffect } from 'react';
import { Table } from 'react-bootstrap';
import { getAllTrackingPlans, getTrackingPlan } from "../../api/trackingplans"
import EventRulesOffCanvas from './EventRulesOffCanvas';
import Pagination from '../Pagination';

const TrackingPlanTable = () => {
  const [trackingPlans, setTrackingPlans] = useState([]);
  const [selectedPlanId, setSelectedPlanId] = useState(null);
  const [trackingPlanDetail, setTrackingPlanDetail] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(-1);

  const [plansPerPage] = useState(5);

  useEffect(() => {
  	fetchTrackingPlans();
  }, [currentPage]);

  const fetchTrackingPlans = async () => {
  	let trackingplans = await getAllTrackingPlans(currentPage, plansPerPage)
  	const tp = Math.ceil(trackingplans.pagination.total / plansPerPage);
  	if (totalPages === -1)
  		setTotalPages(tp)
  	setTrackingPlans(trackingplans.items)
  };

  const handleShowEventRules = async (id) => {
  	setSelectedPlanId(id);
  	const eventRules = await getTrackingPlan(id);
  	setTrackingPlanDetail(eventRules.rules)
  };

  const onHideEventRulesOffCanvas = () => {
  	setSelectedPlanId(null);
  	setTrackingPlanDetail([]);
  }

  return (
	<>
	  <Table>
		<thead>
		  <tr>
			<th>S no.</th>
			<th>Display Name</th>
			<th>Description</th>
		  </tr>
		</thead>
		<tbody>
		  {trackingPlans.map((plan, i) => (
			<tr key={plan.id} onClick={() => handleShowEventRules(plan.id)} className="tableRowPointer">
			  <td>{(currentPage-1)*plansPerPage + i+1}</td>
			  <td>{plan.display_name}</td>
			  <td>{plan.description}</td>
			</tr>
		  ))}
		</tbody>
	  </Table>
	  <Pagination
		currentPage={currentPage}
		totalPages={totalPages}
		onChangePage={setCurrentPage}
	  />
	  {selectedPlanId && (
		<EventRulesOffCanvas
		  show={true}
		  onHide={onHideEventRulesOffCanvas}
		  trackingPlanDetail={trackingPlanDetail}
		/>
	  )}
	</>
  );
};

export default TrackingPlanTable;
