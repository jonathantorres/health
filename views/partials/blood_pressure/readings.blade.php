<div class="panel panel-default">
    <div class="panel-heading">{{ $title }}</div>
    <div class="panel-body">
        @if (count($bloodPressureReadings) > 0)
            <div class="table-responsive">
                <table class="table table-striped table-hover table-condensed">
                    <thead>
                        <tr>
                            <th class="text-center">SYS</th>
                            <th class="text-center">DIA</th>
                            <th class="text-center">Pulse</th>
                            <th class="text-center">Date</th>
                            <th class="text-center">Severity</th>
                            <th class="text-center">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach($bloodPressureReadings as $reading)
                            <tr>
                                <td class="text-center">{{ $reading->systolic }}</td>
                                <td class="text-center">{{ $reading->diastolic }}</td>
                                <td class="text-center">{{ $reading->pulse }}</td>
                                <td class="text-center">{{ date('M, j Y', strtotime($reading->reading_date)) }}</td>
                                <td class="text-center text-{{ $reading->severity()['class'] }}">
                                    {{ $reading->severity()['text'] }}
                                </td>
                                <td class="text-center">
                                    <a href="{{ route('blood-pressure-details', ['id' => $reading->id]) }}" data-toggle="tooltip" data-placement="top" title="View reading details">
                                        <span class="glyphicon glyphicon-search" aria-hidden="true"></span>
                                    </a>
                                    <a href="{{ route('edit-blood-pressure', ['id' => $reading->id]) }}" data-toggle="tooltip" data-placement="top" title="Edit reading">
                                        <span class="glyphicon glyphicon-edit" aria-hidden="true"></span>
                                    </a>
                                    <a href="{{ route('delete-blood-pressure', ['id' => $reading->id]) }}" data-toggle="tooltip" data-placement="top" title="Delete reading" data-confirm="confirm" data-message="Are you sure you wish to delete this reading?">
                                        <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                                    </a>
                                </td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        @else
            <div class="alert alert-warning">
                There are no blood pressure readings.
            </div>
        @endif
    </div>
</div>
