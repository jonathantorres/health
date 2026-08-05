<?php

namespace App;

use Illuminate\Auth\Passwords\CanResetPassword;
use Illuminate\Contracts\Auth\CanResetPassword as CanResetPasswordContract;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;

class User extends Authenticatable implements CanResetPasswordContract
{
    use HasFactory, Notifiable, SoftDeletes, CanResetPassword;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'last_name', 'email', 'password',
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token',
    ];

    /**
     * The attributes that should be cast.
     *
     * @var array
     */
    protected $casts = [
        'deleted_at' => 'datetime',
    ];

    /**
     * The blood pressures entered by the user.
     *
     * @return HasMany
     */
    public function bloodPressures()
    {
        return $this->hasMany(BloodPressure::class);
    }

    /**
     * The weight entries entered by the user.
     *
     * @return HasMany
     */
    public function weights()
    {
        return $this->hasMany(Weight::class);
    }
}
