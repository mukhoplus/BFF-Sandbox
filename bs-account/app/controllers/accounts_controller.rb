class AccountsController < ApplicationController
  require 'json'

  def show
    data = load_data
    id = params[:id].to_i
    account = data.find { |acc| acc["id"] == id }

    if account
      render json: account
    else
      render json: { error: "Account with id #{id} not found" }, status: :not_found
    end
  end

  def by_customer
    data = load_data
    customer_id = params[:customer_id].to_i
    accounts = data.select { |acc| acc["customer_id"] == customer_id }

    if accounts.any?
      render json: accounts
    else
      render json: { error: "No accounts found for customer_id #{customer_id}" }, status: :not_found
    end
  end

  private

  def load_data
    file_path = Rails.root.join('db.json')
    file = File.read(file_path)
    JSON.parse(file)
  end
end
