<div class="panel panel-default">
    <div class="panel-heading">{{ $title }}</div>
    <div class="panel-body">
        @if (count($weightEntries) > 0)
            <div class="table-responsive">
                <table class="table table-striped table-hover table-condensed">
                    <thead>
                        <tr>
                            <th class="text-center">Weight</th>
                            <th class="text-center">Date</th>
                            <th class="text-center">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach($weightEntries as $entry)
                            <tr>
                                <td class="text-center">{{ number_format($entry->weight, 1) }} lbs</td>
                                <td class="text-center">{{ date('M, j Y', strtotime($entry->entered_date)) }}</td>
                                <td class="text-center">
                                    <a href="{{ route('edit-weight', ['id' => $entry->id]) }}">
                                        <span class="glyphicon glyphicon-edit" aria-hidden="true"></span>
                                    </a>
                                </td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            </div>
        @else
            <div class="alert alert-warning">
                There are no weight entries.
            </div>
        @endif
    </div>
</div>
