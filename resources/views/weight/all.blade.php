@extends('layouts.app')

@section('content')
@include('partials.nav')

<div class="container">
    <div class="row">
        <div class="col-md-12">
            @include('partials.weight.entries')
        </div>
    </div>
    <div class="row">
        <div class="col-md-12 text-center">
            {{ $weightEntries->links() }}
        </div>
    </div>
    @include('partials.footer')
</div>
@endsection
