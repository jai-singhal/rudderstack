import React, { useState, useEffect } from 'react';
import { Table, Button } from 'react-bootstrap';
import { getAllTrackingPlans, getTrackingPlan } from "../../api/trackingplans"
import EventRulesOffCanvas from './EventRulesOffCanvas';
import Pagination from '../Pagination';
import TrackingPlanModal from './TrackingPlanModal'

const TrackingPlanTable = () => {
	const [trackingPlans, setTrackingPlans] = useState([]);
	const [selectedPlanId, setSelectedPlanId] = useState(null);
	const [trackingPlanDetail, setTrackingPlanDetail] = useState([]);
	const [currentPage, setCurrentPage] = useState(1);
	const [totalPages, setTotalPages] = useState(-1);
	const [showTrackingPlanModalArr, setShowTrackingPlanModalArr] = useState(trackingPlans.map(() => false));

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
		setTrackingPlanDetail(eventRules.events)
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
			<th>Options</th>
		  </tr>
		</thead>
		<tbody>
		{trackingPlans.map((plan, i) => (
			<tr key={plan.id}  className="tableRowPointer">
				<td onClick={() => handleShowEventRules(plan.id)}>{(currentPage-1)*plansPerPage + i+1}</td>
				<td onClick={() => handleShowEventRules(plan.id)}>{plan.display_name}</td>
				<td onClick={() => handleShowEventRules(plan.id)}>{plan.description}</td>
				<td>
				<Button onClick={() => {
					const newShowTrackingPlanModalArr = [...showTrackingPlanModalArr];
					newShowTrackingPlanModalArr[i] = true;
					setShowTrackingPlanModalArr(newShowTrackingPlanModalArr);
				}}>Update</Button>
				<TrackingPlanModal
					show={showTrackingPlanModalArr[i]}
					onHide={() => {
					const newShowTrackingPlanModalArr = [...showTrackingPlanModalArr];
					newShowTrackingPlanModalArr[i] = false;
					setShowTrackingPlanModalArr(newShowTrackingPlanModalArr);
					}}
					onSubmit={(data) => {console.log(data);}}
					isUpdate = {true} 
					trackingPlanId = {plan.id}
				/>
				</td>
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
