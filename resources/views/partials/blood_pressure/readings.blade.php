<div class="panel panel-default">
    <div class="panel-heading">Blood Pressure Readings</div>
    <div class="panel-body">
        <div class="table-responsive">
            <table class="table table-striped table-hover table-condensed">
                <thead>
                    <tr>
                        <th class="text-center">SYS</th>
                        <th class="text-center">DIA</th>
                        <th class="text-center">Pulse</th>
                        <th class="text-center">Date</th>
                        <th class="text-center">Severity</th>
                        <th class="text-center">View Details</th>
                    </tr>
                </thead>
                <tbody>
                    @foreach($readings as $reading)
                        <tr>
                            <td class="text-center">{{ $reading->systolic }}</td>
                            <td class="text-center">{{ $reading->diastolic }}</td>
                            <td class="text-center">{{ $reading->pulse }}</td>
                            <td class="text-center">{{ date('M, j Y', strtotime($reading->reading_date)) }}</td>
                            <td class="text-center">{{ $reading->severity() }}</td>
                            <td class="text-center">
                                <a href="{{ route('blood-pressure-details', ['id' => $reading->id]) }}">
                                    <span class="glyphicon glyphicon-search" aria-hidden="true"></span>
                                </a>
                            </td>
                        </tr>
                    @endforeach
                </tbody>
            </table>
        </div>
    </div>
</div>
