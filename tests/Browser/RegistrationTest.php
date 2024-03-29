<?php

namespace Tests\Browser;

use App\User;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Laravel\Dusk\Browser;
use Tests\DuskTestCase;

class RegistrationTest extends DuskTestCase
{
    use DatabaseMigrations;

    public function setUp()
    {
        parent::setUp();
    }

    /** @test */
    public function register_guest_user()
    {
        $this->browse(function (Browser $browser) {
            $browser->visit('/register')
                    ->assertSee('Register')
                    ->type('name', 'Charles')
                    ->type('last_name', 'Bronson')
                    ->type('email', 'charles@strong.com')
                    ->type('password', 'strong')
                    ->type('password_confirmation', 'strong')
                    ->press('Register')
                    ->assertPathIs('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->logout();
        });
    }

    /** @test */
    public function existing_user_should_not_be_able_to_register()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->visit('/register')
                    ->assertSee('Register')
                    ->type('name', $user->name)
                    ->type('last_name', $user->last_name)
                    ->type('email', $user->email)
                    ->type('password', 'secret')
                    ->type('password_confirmation', 'secret')
                    ->press('Register')
                    ->assertPathIs('/register')
                    ->assertSee('The email has already been taken.');
        });
    }
}
