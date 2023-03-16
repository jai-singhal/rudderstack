import {React, useEffect} from 'react';
import TrackingPlan from '../components/trackingplan/TrackingPlan';

const Home = () => {
	useEffect(() => {
		document.title = 'Rudderstack Tracking Plan';
	}, []);

  	return (
		<TrackingPlan />
	)
}

export default Home