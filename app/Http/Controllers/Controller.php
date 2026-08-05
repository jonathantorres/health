<?php

namespace App\Http\Controllers;

use App\Version;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Routing\Controller as BaseController;

class Controller extends BaseController
{
    use AuthorizesRequests;

    /**
     * The data that is going to be passed to the view.
     *
     * @var array
     */
    protected $data = [];

    public function __construct()
    {
        $this->data['version'] = Version::VERSION;
    }
}
